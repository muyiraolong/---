package main

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMSG(i int) {
	defer wp.Done()
	fmt.Println(i)
}

func main() {
	for i := 0; i < 10; i++ {
		go showMSG(i)
		wp.Add(1)
	}
	wp.Wait()

	// time.Sleep(time.Second)
}
