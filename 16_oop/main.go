package main

import (
	"fmt"
)

type MyInt int

func (mi MyInt) Double() int {
	return int(mi * 2)
}

func main() {
	v := MyInt(12)

	// calling method related to MyInt
	fmt.Println(v.Double())
	/*
	  - here v is the implicit argument to the func Double
	*/
}
