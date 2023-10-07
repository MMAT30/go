package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	adder := adder(1, 2)
	add, sub := adderSubber(1, 2)
	sum := sum(nums...)


	fmt.Println(adder)
	fmt.Println(add, sub)
	fmt.Println(sum)


	tick := ticker()

	tick()
	tick()
	tick()


}

// single return type
func adder(a int, b int) int {
	return a + b
}

// multiple return types
func adderSubber(a int, b int) (int, int) {
	return a + b, a - b
}


// variatic function
func sum(c ...int) int {
	
	sum := 0
	for _, num := range c {
		sum += num
	}

	return sum
}

// closure
func ticker() (func()) {
	sum := 0
	return func() {
		sum += 1
		fmt.Println("tick", sum)
	}
}
