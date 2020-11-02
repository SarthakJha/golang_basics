package main

import "fmt"

func main() {
	// init map
	// age := make(map[string]int)

	// init and assign simultaneosly
	age := map[string]int{
		"sarthak": 18,
		"sanjeev": 88,
		"a":       1,
		"b":       2,
	}

	// assigning val to maps
	age["saransh"] = 13
	// sanjeev's age modified here
	age["sanjeev"] = 47
	age["vibha"] = 45

	fmt.Println(age)
	println(len(age))

	// deleting from map
	delete(age, "sarthak")

	// iterating over a map
	for key, value := range age {
		fmt.Printf("%s is of age %d\n", key, value)
	}

	// println(age) returns the address of age map
	fmt.Println(age)
}
