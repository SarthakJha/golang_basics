package main

import "fmt"

func main() {
	ids := []int{1, 2, 3, 4, 5, 6, 7, 88}

	// loop through ids (with index)
	for i, id := range ids {
		fmt.Printf("index: %d, value: %d\n", i, id)
	}

	// loop through ids without index
	// same could be used to iterate over just keys or just values in maps
	for _, id := range ids {
		fmt.Printf("index:, value: %d\n", id)
	}
}
