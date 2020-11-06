// Using channels in goroutines

package main

import (
  "fmt"
  "sync"
  "time"
)

func main() {

  var wg sync.WaitGroup // declaring a wait group
  out1 := make(chan string)
  out2 := make(chan string)

  wg.Add(1) // increment wg for every goroutine
  go PrintH("hello world", &wg, out1)

  wg.Add(1)
  go PrintH("#2 Hello world", &wg, out2)
  
/*
one to handle messages from channels

  for {
    msg1, open1 := <-out1 // if the channel is closed open1 will be false
    msg2, open2 := <-out2 // similarly open2 will be false
    if !open1 || !open2 {
      break
    }
    fmt.Println(msg1,msg2)
  } 
 */
 
 // processes output of the routine which finishes first (unlike sequential for loops)
 for {
   select {
    case msg := <-out1:
      fmt.Println(msg)
    case msg := <-out2:
      fmt.Println(msg)
   }
 }

  wg.Wait() // this will wait wg count is zero i.e all goroutines have finished and then end the main thread

  // fmt.Scanln() // one way to handle routine exit
}

// out chan string is the declaration of the channel out which transports string
func PrintH(str string, wg *sync.WaitGroup, out chan string) {
  // NEVER CLOSE THE CHANNEL FROM RECIEVING END (panic error)
  defer close(out) // this closes the channel before exiting function to avoid panic error
  for i:=0;i <= 9;i++ {
			time.Sleep(time.Second)
  		out <- str
		}	
    wg.Done() // decrement the wg count i.e mark the end of a goroutine
}
