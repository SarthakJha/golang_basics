/*
VVVVVVV.IMP
- defer calls are evaluated immediately
- BUT are called after execution of rest of the main
*/

package main

import "fmt"

func main() {
	i := 1
	defer fmt.Println(i + 1) // 2 (expected 3 but immediately evaled)
	i++
  
	/*
		- this ex shows defer calls are evaluated immediately
			but are called at the end of the main function
	*/
}
