package main

import (
	"fmt"
	"math"
)

// closure example
func makeDistOrigin(o_x, o_y float64) func(float64, float64) float64 {
	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}
	return fn
}

// function that takes function as argument
func applyIt(aFunct func(int) int, val int) int {
	return aFunct(val)
}

// function to be added as parameter
func incrVal(val int) int {
	return val + 1
}
func decrVal(val int) int {
	return val - 1
}

func main() {
	// takes in origin, returns a function to calc distance from custom origin
	fn := makeDistOrigin(0, 0)
	// takes in point to calc dist from origin, returns distance
	dist := fn(2, 3)

	// calling the function
	incrementedVal := applyIt(incrVal, 2)
	decrementedVal := applyIt(decrVal, 2)

	fmt.Println(incrementedVal, decrementedVal)
	fmt.Printf("the distance from custom origin is %f\n", dist)

}
