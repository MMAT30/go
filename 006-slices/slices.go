package main

import "fmt"

func main() {

	a1 := make([]int, 3)
	a2 := []int{7,7,7}
	a3 := []int{1,2,3,4,5, 6, 7, 8, 9, 10}

	// copying a1 into a2
	fmt.Println("a1 before:", a1)
	fmt.Println("a2 before:", a2)
	fmt.Println("copying a1 into a2")

	copy(a1, a2)

	fmt.Println("a4 after:", a1)
	fmt.Println("a5 after:", a2)
	fmt.Println()

	// appending to a slice
	fmt.Println("a1 before:", a1)
	fmt.Println("appending 8 to a1")

	a1 = append(a1, 8)

	fmt.Println("a1 after:", a1)
	fmt.Println()


	// slicing arrays
	fmt.Println("a3[:]:", a3[:])
	fmt.Println("a3[:5]:", a3[:5])
	fmt.Println("a3[5:]:", a3[5:])
	fmt.Println("a3[len(a3) - 1:]:", a3[len(a3) - 1:])
	fmt.Println("a3[len(a3) - 3:]:", a3[len(a3) - 3:])
	fmt.Println("a3[:len(a3) - 1]:", a3[:len(a3) - 1])
	fmt.Println("a3[:len(a3) - 3]:", a3[:len(a3) - 3])


	

}