package main

import (
	"fmt"
	"math"
)

func genDisplaceFn(a, u, s float64) func() float64 {
	v := math.Sqrt(math.Pow(u, 2) + (2 * a * s))
	return func() float64 {
		t := (v - u) / a
		return t
	}
}
func main() {
	fn := genDisplaceFn(2, 0, 10)
	fmt.Println(fn())
}
