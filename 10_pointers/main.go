package main

import "fmt"

func main() {
	var a *int
	b := 10
	a = &b

	fmt.Printf("value: %d, addr: %p\n", *a, a)
}
