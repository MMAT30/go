package main

import (
	"fmt"
	"reflect"
)

func main() {

	m1 := map[int]string{1:"one", 2:"two", 3:"three"}

	// type and data
	fmt.Println(m1, reflect.TypeOf(m1))

	// adding elements
	m1[4] = "four"
	fmt.Println(m1, reflect.TypeOf(m1))

	// deleting from map
	delete(m1, 4)
	fmt.Println(m1, reflect.TypeOf(m1))


}