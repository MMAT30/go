package main

import (
	"fmt"
	"os"
)

func main() {
	
	programName := os.Args[0]
	args := os.Args[1:]
	
	fmt.Println("Program Name:", programName)
	fmt.Println("Arguments:", args)

}