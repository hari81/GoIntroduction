package main

import (
	"fmt"
	"math"
)

// Define interface

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}
type Reactange struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Reactange) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	fmt.Println("Interfaces:")
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Reactange{width: 10, height: 5}
	fmt.Printf("Circle Area %f\n", getArea(circle))
	fmt.Printf("Rectangle Area %f\n", getArea(rectangle))
}
