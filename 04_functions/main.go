package main

import (
	"fmt"
)

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Sarthak"))
	fmt.Println("sum of 3 and 7 is " + fmt.Sprint(getSum(3, 7)))
}

func greeting(name string) string {
	return "hello, " + name + "!"
}
