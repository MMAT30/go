package main

import (
	"fmt"
)

func main() {

	var v1 int = 1
	var v2 float64 = 1.0
	var v3 byte = 0x01
	var v4 rune = 0x01
	var v5 bool = true
	var v6 string = "hello world"
	var v7 []int = []int{1, 2, 3}
	var v8 map[string]int = map[string]int{"a": 1, "b": 2, "c": 3}

	var v9 struct {
		x int
		y int
		z int
	} = struct {
		x int
		y int
		z int
	}{1, 2, 3}

	


	fmt.Println("int:\t\t", v1)
	fmt.Println("float64:\t", v2)
	fmt.Println("byte:\t\t", v3)
	fmt.Println("rune:\t\t", v4)
	fmt.Println("bool:\t\t", v5)
	fmt.Println("string:\t\t", v6)
	fmt.Println("array:\t\t", v7)
	fmt.Println("map:\t\t", v8)
	fmt.Println("struct:\t\t", v9)
}