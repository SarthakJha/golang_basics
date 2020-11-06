/*
* working with channels and goroutines
* */
package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
// another way to handle rountines is waitGroup (sync.WaitGroup)
  var wg sync.WaitGroup // declaring a wait group
/*
* IF YOU DONT WANT TO USE POINTERS to send wg, use anonymous func

go func() {
  PrintH("hello world")
  wg.Done()
}()

* */
  wg.Add(1) // increment wg for every goroutine
  go PrintH("hello world", &wg)

  wg.Add(1)
  go PrintH("#2 Hello world", &wg)


  wg.Wait() // this will wait wg count is zero i.e all goroutines have finished and then end the main thread

  //  fmt.Scanln() // one way to handle routine exit
}

func PrintH(str string, wg *sync.WaitGroup) {
  for i:=0;i <= 9;i++ {
			fmt.Println(str,i)
			time.Sleep(time.Second)
		}	
    wg.Done() // decrement the wg count i.e mark the end of a goroutine
}
