package main

import (
	"fmt"
	// "reflect"
)

// non returning functions
func f1(str string) {
	fmt.Println(str)
}

func f2(str1 string, str2 string) (string, string) {
	return str1, str2
}

func f3(nums ...int) int{
	sum := 0

	for num := range nums {
		sum += num
	}

	return sum
}

func f4() (func() string) {
	return func () string {
		return "f4 function"
	}
}


func f6() (func() int) {
	cnt := 0

	return func () int {
		cnt++
		return cnt
	}
}


func sum(slice ...int) int {
	sum := 0

	for x := range slice {
		sum += x
	}
	return sum
}


func even(f func(slice ...int) int, nums ...int) bool {

	sum := f(nums...)

	if (sum % 2 == 0) {
		return true
	} else {
		return false
	}
}


func main() {

	a1 := []int{1,2,3,4,5}

	// defering function f1 till the end of main
	defer f1("f1 function")

	// calling functions
	fmt.Println(f2("f2", "function"))
	fmt.Println(f3(a1...))

	// function return a function
	innerFunc := f4()
	fmt.Println(innerFunc())

	// function expressions
	f5 := func (str string)  {
		fmt.Println(str)
	}
	f5("f5 function expression")


	// closure
	adder := f6()
	fmt.Println(adder())
	fmt.Println(adder())
	fmt.Println(adder())
	fmt.Println(adder())


	// callback
	fmt.Println(even(sum, []int{1,2,3,4,5}...))


	
}
