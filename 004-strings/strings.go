package main

import "fmt"

func main() {

	str := "taco!"
	fmt.Println("string:", str)



	// string concatenation
	fmt.Println("concatenation:", str + " cat")

	// string length
	fmt.Println("length:", len(str))

	// rune count
	fmt.Println("rune count:", len([]rune(str)))

	// string as array of bytes
	fmt.Println("strings as bytes:", []byte(str))

	// string as array of runes
	fmt.Println("strings as runes:", []rune(str))


	// looping over string
	for index, val := range str {
		fmt.Println(index, string(val), byte(val), rune(val))
	}

	
}