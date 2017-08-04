package main

import (
	"flag"
	"io/ioutil"
	"strings"

	"github.com/seblegall/parallel"
)

func main() {

	var file string
	var workers int

	flag.StringVar(&file, "f", "", `The file containing the commands to play`)
	flag.IntVar(&workers, "w", 2, `The number of parallel processus to use`)
	flag.Parse()

	fileContent, _ := ioutil.ReadFile(file)
	commands := strings.Split(string(fileContent), "\n")
	process := parallel.NewParallel(commands, workers)
	process.Launch()
	process.Print()
}
