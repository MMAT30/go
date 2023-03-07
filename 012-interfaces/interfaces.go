package main

import (
	"fmt"
	"math"
)


// shapes interface
type shapes interface {
	area() float64
}




// circle struct
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}



// rectangle struct
type rectangle struct {
	width float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}



// function based on shapes interface
func getArea(s shapes) float64 {
	return s.area()
}





func main() {

	// creating rectangle and circle structs
	r1 := rectangle{width: 5.0, height: 5.0}
	c1 := circle{radius: 5.0}


	// interface is used to group the shapes
	s1 := []shapes{r1, c1}


	// looping thru shapes structs
	for index, obj := range s1 {
		fmt.Println(index, getArea(obj))
	}

	
}