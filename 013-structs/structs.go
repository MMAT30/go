package main

import "fmt"

func main() {

	p1 := newPerson("John", 30)
	p2 := newPerson("Jane", 25)

	fmt.Println(p1)
	fmt.Println(p2)


}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = age
	return &p
}


type person struct {
	name string
	age int

}