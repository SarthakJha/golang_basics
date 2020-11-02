package main

import (
	"fmt"
)

// ":=" type of global declaration cant be done outside a fn
const baapKaNaam = "Sanjeev jha"

func main() {
	var name string = "sarthak jha"
	// shorthand declaration
	age := 18
	height := 12.4452
	const married = false
	age = 36
	var heightInMeter float32 = 2.6
	var char rune = 'j' // 106

	// multiple shorthand declaration
	school, code := "MIT", 110088

	// printing out variables
	fmt.Println("my name is " + name)
	fmt.Println(age)
	fmt.Println(married)
	fmt.Println(height)
	fmt.Println(baapKaNaam)
	fmt.Println(heightInMeter)
	fmt.Println(char)
	fmt.Println(school, code)

	// printing out types of the variables
	fmt.Printf("type of name: %T\n", name)
	fmt.Printf("type of age: %T\n", age)
	fmt.Printf("type of married: %T\n", married)
	fmt.Printf("type of height: %T\n", height)
	fmt.Printf("type of char: %T\n", char)

}
