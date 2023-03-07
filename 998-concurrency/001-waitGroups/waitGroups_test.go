package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrinter(t *testing.T) {

	// creating a wait group
	var wg sync.WaitGroup


	// saving default std out
	stdOut := os.Stdout

	// creating a pipe to catch output from std out
	read, write, _ := os.Pipe()
	
	// writing to std out via the pipe
	os.Stdout = write

	// adding to wait group
	wg.Add(1)




	// creating go routine
	go printer("test go routine", &wg)

	// wait for go routine and closing the pipe
	wg.Wait()
	write.Close()

	// reading what was written to the pipe
	result, _ := io.ReadAll(read)
	// converting to string
	output := string(result)

	// resetting std out back to normal
	os.Stdout = stdOut



	// checking if output is correct
	if !strings.Contains(output, "test go routine") {
		t.Errorf("expected to find string 'test go routine'")
	}
}
