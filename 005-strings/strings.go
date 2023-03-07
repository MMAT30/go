package main

import (
	"fmt"
	"reflect"
)


func main() {

	// creating string
	str1 := "string"

	// creating charater or rune string of size 10 with a max capacity of 15
	str2 := string(make([]rune, 10, 15))


	fmt.Println(reflect.TypeOf(str1), reflect.TypeOf(str2))
	fmt.Println(len(str1), len(str2))


	for index, char := range str1 {
		fmt.Println(index ,string(char))
	}
}