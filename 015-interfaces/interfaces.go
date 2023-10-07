package main

import (
	"fmt"
)

func main() {
	s := savings{0}
	c := checkings{0}
	fmt.Println("initial savings:", s.getSavings())
	fmt.Println("initial checking:", c.getCheckings())
	fmt.Println()


	s.deposit(100)
	fmt.Println("savings after deposit:", s.getSavings())

	c.deposit(100)
	fmt.Println("checking after deposit:", c.getCheckings())
	fmt.Println()




	s.withdrawl(10)
	fmt.Println("savings after withdrawl:", s.getSavings())

	c.withdrawl(10)
	fmt.Println("checking after withdrawl:", c.getCheckings())
	fmt.Println()





	fmt.Println("savings balance:", balances(&s))
	

}






type accounts interface {
	getCheckings() float64
	getSavings() float64
}

type savings struct {
	account float64	
}

type checkings struct {
	account float64
}





func (s *savings) deposit(amount float64) {
	s.account += amount
}

func (s *savings) withdrawl(amount float64) {
	if (s.account - amount) > 0 {
		s.account -= amount
	}
}

func (s *savings) getSavings() float64 {
	return s.account
}





func (c *checkings) deposit(amount float64) {
	c.account += amount
}

func (c *checkings) withdrawl(amount float64) {
	if c.account - amount > 0 {
		c.account -= amount
	}
}

func (c *checkings) getCheckings() float64 {
	return c.account
}


func balances(a &accounts) {
	fmt.Printf("account type: %T\n", a)

	switch t := a.(type) {
	case *savings:
		fmt.Print("savings balance:", t.getSavings())
	case *checkings:
		fmt.Print("checking balance:", t.getCheckings())

	default:
		fmt.Println("unknown type")
		
	}
}

