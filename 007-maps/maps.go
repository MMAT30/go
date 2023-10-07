package main

import "fmt"

func main() {
	m1 := make(map[string]int)
	fmt.Println(m1)


	m1["a"] = 1
	m1["b"] = 2
	m1["c"] = 3

	fmt.Println(m1)
	delete(m1, "a")

	fmt.Println(m1)
	delete(m1, "b")

	fmt.Println(m1)
	delete(m1, "c")

	
	fmt.Println(m1)
}