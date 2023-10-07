package main 

import "fmt"

func main() {

	var a1 [4]int
	var a2 [4][4]int
	var a3 [4][4][4]int
	count := 1

	// single dimensional array
	for x := 0; x < 3; x++ {
		a1[x] = x
		count++
	}

	// resetting counter
	count = 1

	// 2 dimensional array
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			a2[x][y] = count
			count++
		}
	}


	// resetting counter
	count = 1


	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			for z := 0; z < 4; z++ {
				a3[x][y][z] = count
				count++
			}
		}
	}

	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}