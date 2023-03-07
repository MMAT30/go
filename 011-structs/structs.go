package main

import (
	"fmt"
	"reflect"
)



// person struct
type person struct {
	first string
	last string
	age int
}

// worker struct
type worker struct {
	person
	job string
}


// methods
func (w worker) speak() {
	fmt.Sprintf("hello im %s %s and im %s years old", w.first, w.last, w.age)
}


func main() {


	// creating person stucts
	p1 := person{
		first: "john",
		last: "smith",
		age: 777,
	}
	
	p2 := person{
		first: "greg",
		last: "johnson",
		age: 888,
	}

	w1 := worker{
		person: person{
			first: "billy",
			last: "bob",
			age: 999,
		},
		job: "programmer",
	}

	a1 := struct{
		first string
		last string
		age int
	}{
		first: "some",
		last: "one",
		age: 123,
	}

	// anonymous struct
	fmt.Println(a1)


	// person struct and type
	fmt.Println(p1, p2)
	fmt.Println(reflect.TypeOf(p1), reflect.TypeOf(p2))


	// accessing data members
	fmt.Println(p1.first, p1.last, p1.age)


	// accessing embedded struct
	fmt.Println(w1.first, w1.last, w1.age, w1.job)


	// calling methods attched to struct
	w1.speak()

	


}