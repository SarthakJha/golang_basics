package main

import (
	"fmt"
)

func getSum(num1, num2 int) int {
	return num1 + num2
}

// function with multiple return value
func foo2(x int) (int, int) {
  return x, x+1
}
func main() {
	fmt.Println(greeting("Sarthak"))
	fmt.Println("sum of 3 and 7 is " + fmt.Sprint(getSum(3, 7)))
	
	// calling foo2
  x,y:=foo2(4)
  println(x,y) // 4 5
}

func greeting(name string) string {
	return "hello, " + name + "!"
}
