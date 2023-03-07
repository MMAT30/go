package main

import (
	"math/rand"
	"fmt"
	"testing"
	"time"
)


const (
	Pizzas = 10
)
var pizzasMade, pizzasFailed, total int



type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <- ch
}


type PizzaOrder struct {
	PizzaOrder int
	message string
	success bool
}

func pizzeria(pizzaMaker *Producer) {

}


func Test(t *testing.T) {

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// print message
	fmt.Println("pizzas are being made")

	// creating a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}


	// run producer in the background
	go  pizzeria(pizzaJob)


	// create and run consumer
	// print out message

}