package main

import (
	"fmt"
)

func main() {


	a := false
	b := true

	// basics
	fmt.Println(!a)
	fmt.Println(a == a)
	fmt.Println(a != b)
	fmt.Println(a || b)
	fmt.Println(a && b)
	

	// if - else if - else
	if b == true {
		fmt.Println("if statement")
	} else if b == false {
		fmt.Println("else if statement")
	} else {
		fmt.Println("else statement")
	}

	// switch - case - default
	switch b {
	case true:
		fmt.Println("true case")
	case false:
		fmt.Println("false case")
	default:
		fmt.Println("default case")
	}
}
