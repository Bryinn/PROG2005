/*
Create a programme that has a Geometry interface containing an area() method (Signature: "func area() float64").
Create subtypes Circle and Square that each implement the area() method (and calculate the respective area correctly)
Create a function with the signature "func PrintArea(g Geometry)" to which you can pass a Circle or Square instance that prints the correct area (depending on type).
*/
package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
}

// structs
type Circle struct {
	radius float64
}
type Square struct {
	side float64
}

func (c Circle) area() float64 {
	area := c.radius * c.radius * math.Pi
	return area
}

func (s Square) area() float64 {
	area := s.side * s.side
	return area
}

func PrintArea(g Geometry) {
	fmt.Printf("Area: %.2f\n", g.area())
}

func main() {
	c := Circle{radius: 5}
	s := Square{side: 4}

	PrintArea(c)
	PrintArea(s)
}
