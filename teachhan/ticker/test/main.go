package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

// func send()  {
// 	<-5
// }

func main() {
	go func() {
		ch <- 5
		close(ch)
	}()
	fmt.Println(<-ch)

	time.Sleep(time.Second)
}
