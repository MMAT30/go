package main

import (
	"fmt"
)

func main() {
	a := account{0, 0}
	fmt.Println("initial account:", a)
	fmt.Println()


	// we add 100 to the checking account by value
	a.addToCheckingByValue(100)
	// but the value of a.checking is still 0 
	// because we are only passing a copy of the account
	fmt.Println(a)


	// we add 100 to the checking account by reference
	a.addToCheckingByReference(100)
	// now the value of a.checking is 100
	// because we are passing a pointer to the actuall account
	fmt.Println(a)

	


}

type account struct {
	savings, checking int
}

func (a *account) addToCheckingByReference(amount int) {
	fmt.Printf("account type: %T\n", a)
	a.checking += amount
}

func (a account) addToCheckingByValue(amount int) {
	fmt.Printf("account type: %T\n", a)
	a.checking += amount
}
