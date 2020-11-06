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

// using Shape as a type in args
func getArea(s Shape) float64 {
	return s.getArea()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 1.4}
	rectangle := Rectangle{23.2, 3.4}
  
  // accessing getArea() which take shape as arg
 fmt.Println(getArea(circle))

  /*
  * declaring slice of type Shapes
  * legal only of the struct have the method of the interface defined
    and with exactly same args
  * using "shapes", one can only access the "getArea()" fn out of them
    and not height, radius, ect
  * */
  shapes := []Shape{circle, rectangle}

	// printing area of objects in "shapes" slice by for loop
 for _, val := range shapes {
    fmt.Println(val.getArea())
  }
}
