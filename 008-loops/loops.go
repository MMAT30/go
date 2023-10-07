package main

import "fmt"

func main() {
	
	count := 3
	start := 0
	a := []int{1, 2, 3}
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	s := "taco"

	// for loop
	for i := 0; i < count; i++ {
		fmt.Println("for loop:", i)
	}

	// for loop as while loop
	for start < 3 {
		fmt.Println("for loop as while loop:", start)
		start++
	}

	// for range loop array
	for i, v := range a {
		fmt.Println("for range array:", i, ":", v)
	}

	// for range loop key-pair values
	for k, v := range m {
		fmt.Println("for range key-value:", k, ":", v)
	}


	// for range loop string
	for i, v := range s {
		fmt.Println("for range string:", i, ":", v, ":", string(v))
	}
}