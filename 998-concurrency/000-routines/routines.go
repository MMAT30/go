package main

import (
	"fmt"
	"runtime"
)


func printer(str string) {
	fmt.Println(str)
}


func main() {

	// arch data
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("CUPs", runtime.NumCPU())
	fmt.Println()

	fmt.Println("ROUTINES", runtime.NumGoroutine())
	fmt.Println()



	for _, value := range []int{1,2,4,5,6} {

		go printer(fmt.Sprintf("go routine %d", value))
	}

	// we launched 6 go routines but they all die
	// because main exits faster than the routines 
	// can execute


	fmt.Println("ROUTINES", runtime.NumGoroutine())

}