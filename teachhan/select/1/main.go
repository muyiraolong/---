package main

import (
	"fmt"
	"runtime"
)

func main() {
	var chanInt = make(chan int)
	var chanStr = make(chan string)
	go func() {
		chanInt <- 8
		chanStr <- "kk"
		defer close(chanInt)
		defer close(chanStr)
	}()
	for i := 0; i < 5; i++ {
		select {
		case a := <-chanInt:
			fmt.Println("chanInt:", a)
		case b := <-chanStr:
			fmt.Println("chanStr:", b)
		default:
			fmt.Println("Null")
		}
		runtime.Gosched()
	}

}
