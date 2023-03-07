package main

import (
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateString(s string) {
	defer wg.Done()
	msg = s
}

func main() {

	msg = "hello world"

}
