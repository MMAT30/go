package main

import "fmt"

func main() {

	// if-elseif-else
	if true {
		fmt.Println("if statement")
	} else if true != false {
		fmt.Println("else if statement")
	} else {
		fmt.Println("else statement")
	}

	// switch-case-default
	switch "test" {
	case "test":
		fmt.Println("switch statement")
	default:
		fmt.Println("default statement")
	}

	// TODO: type switch
	var t interface{}

	switch t := t.(type) {
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	}

}
