package main

import (
	"fmt"
	"reflect"
)


func main() {

	b1 := true
	b2 := false

	fmt.Println(reflect.TypeOf(b1), reflect.TypeOf(b2))
	fmt.Println(b1 || b2)

}