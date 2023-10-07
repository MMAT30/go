package main

import "fmt"

func main() {

	i := 0
	fmt.Println("i:", i, "address:", &i)

	// Pass the address of i to adder
	adder(&i)
	fmt.Println("i:", i, "address:", &i)

}

func adder(num *int) {

	// address of num
	fmt.Println("num:", num)
	// dereference num and increment the value
	*num++
}

