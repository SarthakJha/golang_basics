package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	getArea() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// Area method for circle
func (c Circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

// Area method for rectangle
func (r Rectangle) getArea() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 1.4}
	rectangle := Rectangle{23.2, 3.4}

	// printing area
	fmt.Printf("area of circle is: %f\n ", getArea(circle))
	fmt.Printf("area of rectangle is: %f\n ", getArea(rectangle))

}
