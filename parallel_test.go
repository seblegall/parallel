package parallel

import (
	"os"
	"os/exec"
	"reflect"
	"sync"
	"testing"
)

type parallelArgs struct {
	cmdList []string
	workers int
}

var testNewParallel = []parallelArgs{
	{[]string{"test1", "test2"}, 0},
	{[]string{"test1", "test2"}, 3},
	{[]string{}, 3},
}

func TestNewParallel(t *testing.T) {

	wd, _ := os.Getwd()

	for _, pArgs := range testNewParallel {
		expected := Parallel{
			workers: pArgs.workers,
			cmdList: pArgs.cmdList,
			wg:      new(sync.WaitGroup),
			cmd:     make(chan *exec.Cmd, 2*pArgs.workers),
			wd:      wd,
		}

		actual := NewParallel(pArgs.cmdList, pArgs.workers)

		if reflect.DeepEqual(expected, actual) {
			t.Fatalf("Construction failed. Expected %+v, get %+v", expected, actual)
		}
	}

}
