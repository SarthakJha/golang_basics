package main

import (
	"fmt"
	"strconv"
)

// Person is ...
type Person struct {
	firstName  string
	secondName string
	city       string
	gender     string
	age        int

	// >> another way of declaring above
	// firstname secondname city gender string
	// age int
}

// greeting method (value reciever)
func (p Person) greet() string {
	p.age++ // this doesnt modify the value, just displays incemented value
	return "Top of the morning to ya, " + p.firstName + " my age: " + strconv.Itoa(p.age)
}

// isMarried method (pointer reciever) used to "modify" value
func (p *Person) hasBirthday() {
	p.age++ // permanenty incemenrts the age of the person
}

// getMarried() (pointer reciever) <change last name>
func (p *Person) isMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.secondName = spouseLastName
	}
}

func main() {
	// init person using struct
	per1 := Person{firstName: "sarthak", secondName: "jha", city: "delhi", gender: "male", age: 19}
	// alternate way of initializing person
	per2 := Person{"karen", "bobby", "boston", "female", 1}

	println(per1.greet())
	println(per2.greet())

	// permanently increments the age
	per1.hasBirthday()
	per2.hasBirthday()
	per2.isMarried(per1.secondName)

	println(per1.age)
	println(per2.age)
	fmt.Println(per1, per2)
}
