package main

import "fmt"

func main() {
	// var arr [2]string
	// arr[0] = "sarthak"
	// arr[1] = "jha"

	// declare and assign(array)
	arr := [2]string{"sarthak", "jha"}

	// declare and assign (slices)
	slice := []string{"vijay", "deenanath", "chauhan"}

	fmt.Println(arr)
	fmt.Println(slice)
	fmt.Printf("Array length is %d\n", len(arr))
	fmt.Printf("slice length is %d\n", len(slice))
	fmt.Printf("type of slice is %T\n", slice)
	fmt.Printf("type of arr is %T\n", arr)

	/*
		- arrays in go have fix length
		- slices dont have a fixed length
	*/
}
