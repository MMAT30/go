package main

import (
	"fmt"
	"reflect"
)

func main() {


	m1 := map[int]string{1:"one", 2:"two", 3:"three"}
	a1 := []int{1,2,3}

	// types
	fmt.Println(reflect.TypeOf(m1), reflect.TypeOf(a1))


	// looping thru map
	for key, value := range m1 {
		fmt.Println(key, value)
	}



	// looping thru array
	for index, value := range a1 {
		fmt.Println(index, value)
	}

	// range 
	for index, value := range []int{1,2,3} {
		fmt.Println(index, value)
	}

	// creating loop
	for x :=0; x <= 4; x++ {
		fmt.Println(x)
	}

}