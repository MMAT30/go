package main

import (
	"fmt"
	"sync"
)


func printer(str string, wg *sync.WaitGroup) {
	// defering completion of the go routine
	defer wg.Done()

	fmt.Println(str)
}


func main() {

	// creating a wait group
	var wg sync.WaitGroup

	// creating array
	nums := []int{1,2,3,4,5,6}


	// adding number of routines to wait for
	wg.Add(len(nums))



	// launching go routines
	for _, value := range nums {
		go printer(fmt.Sprintf("go routine [%d/%d]", value, len(nums) ), &wg)
	}



	// waiting for the go routines to complete
	wg.Wait()

	// then execute the rest of the code
	fmt.Println("Completed")
}