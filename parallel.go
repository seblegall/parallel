package parallel

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"text/tabwriter"
)

type cmdError struct {
	err error
	cmd string
}

//Parallel represents a process that will exec cmds in parallel.
type Parallel struct {
	workers    int
	errorCount int
	errors     []cmdError
	wg         *sync.WaitGroup
	cmd        chan *exec.Cmd
	cmdList    []string
	wd         string
}

//NewParallel init a new reload command.
func NewParallel(cmdList []string, workers int) Parallel {

	//Set default for workers
	if workers == 0 {
		workers = 2
	}

	//get current working dir
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return Parallel{
		workers: workers,
		cmdList: cmdList,
		wg:      new(sync.WaitGroup),
		cmd:     make(chan *exec.Cmd, 2*workers),
		wd:      wd,
	}
}

//Launch actualy execute the reload command
//This function print a progress bar to follow the test current status.
func (p *Parallel) Launch() {

	fmt.Println("Starting parallel processus...")

	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func(p *Parallel) {
			defer p.wg.Done()
			for c := range p.cmd {
				//Start process. Exit code 127 if process fail to start.
				if err := c.Start(); err != nil {
					p.errorCount++
					p.errors = append(p.errors, newCmdError(err, c.Args))
				}

				if err := c.Wait(); err != nil {
					p.errorCount++
					p.errors = append(p.errors, newCmdError(err, c.Args))
				}
			}
		}(p)
	}

	for _, cmd := range p.cmdList {

		c := exec.Command("/bin/sh", "-c", cmd)

		c.Dir = p.wd

		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		p.cmd <- c
	}

	close(p.cmd)
	p.wg.Wait()
}

//Print show the commands result using tab writer.
func (p *Parallel) Print() {

	fmt.Printf(`


Total command in error : %d

`, p.errorCount)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Cmd\tError")
	for _, err := range p.errors {
		fmt.Fprintf(w, fmt.Sprintf("%s\t%s", err.cmd, err.err.Error()))
		fmt.Fprintln(w)
	}
	w.Flush()
}

func newCmdError(err error, args []string) cmdError {
	return cmdError{
		err: err,
		cmd: strings.Join(args, " "),
	}
}
