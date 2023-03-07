package main

import (
	"io"
	"os"
	"strings"
	"testing"
)


func TestMain(t *testing.T) {

	// original state of std out
	stdOut := os.Stdout

	// creating a pipe to redirect std out
	read, write, _ := os.Pipe()


	// std out now writes to the pipe
	os.Stdout = write
	


	// main
	main()





	// closing writing to pipe
	write.Close()

	// reading bytes from pipe and casting to string
	result, _ := io.ReadAll(read)
	output := string(result)

	// resetting std out to original state
	os.Stdout = stdOut


	// checking balance
	if ! strings.Contains(output, "135200") {
		t.Error("wrong balance returned")
	}
}