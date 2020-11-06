package main

import "fmt"

// pass by reference sorting function
// arrays are by default passed as reference
func bubbleSort(x []int) {
	fmt.Printf("value: %d, addr: %p\n", x, &x)
	fmt.Printf("type of x: %T\n", x)
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x)-i-1; j++ {
			fmt.Printf("j is %d\n", j)
			if x[j] > x[j+1] {
				// interchanging values
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}

/*
- if passed as pointer (bubbleSort(x *int[])) then,
- access arrays like (*x)[i]
*/

func main() {
	var a *int
	b := 10
	a = &b

	// declaring slice
	x := []int{1, 6, 3, 7, 5, 2}

	bubbleSort(x)
	fmt.Printf("value: %d, addr: %p\n", *a, a)
	fmt.Printf("value: %d, addr: %p\n", x, &x)
}
