package main

import (
	"fmt"
	"sync"
)

var i int

var w1, w2 sync.WaitGroup
var lock sync.Mutex

func add() {
	defer w1.Done()
	lock.Lock()
	i++
	defer lock.Unlock()
}

func sub() {
	defer w2.Done()
	lock.Lock()
	i--
	defer lock.Unlock()
}

func main() {
	for j := 0; j < 100; j++ {
		w1.Add(1)
		go add()
		w2.Add(1)
		go sub()
	}
	w1.Wait()
	w2.Wait()
	// runtime.Gosched()
	// time.Sleep(time.Second)
	// time.Sleep(time.Second * 3)
	fmt.Println(i)
}
