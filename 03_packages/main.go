package main

import (
	"fmt"
	"math"

	"github.com/SarthakJha/go_tutorial/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.99)) // 2
	fmt.Println(math.Ceil(2.01))  // 3
	fmt.Println(math.Sqrt(5))

	// reverse string function from strutil
	fmt.Println(strutil.Reverse("hello world"))
}
