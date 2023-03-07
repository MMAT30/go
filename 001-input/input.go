package main

import (
	"fmt"
	"os"
)

func main() {


	// arguments
	args := os.Args


	// size
	argc := len(os.Args)


	fmt.Println(args, argc)
	
}