package main

import (
	"fmt"
	"reflect"
)

func main() {

	// creating sized array
	a1 := [4]int{1,2,3,4}

	// creating unsized array
	a2 := []int{1,2,3,4}

	// creating empty sized array with a max capacity
	a3 := make([]int, 10, 20)

	// creating unsized mulitdimensional array
	a4 := [][]int{{1,2,3}, {4,5,6}}

	// creating sized mulitdimensional array
	a5 := [2][3]int{{1,2,3}, {4,5,6}}

	// data and types
	fmt.Println(a1, a2, a3, a4, a5)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(a2), reflect.TypeOf(a3), reflect.TypeOf(a4), reflect.TypeOf(a5))



	// accessing the array
	fmt.Println(a1[0])
	
	// appending to the array
	a2 = append(a2, 5,6)
	fmt.Println(a2)

	// deleting 3rd element by using the append function and unfurling the array slice
	a2 = append( a2[:2], a2[4:]...)
}