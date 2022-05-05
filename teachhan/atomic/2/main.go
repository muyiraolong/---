// 原理：compare and swap
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

var i int32

func add() {
	atomic.AddInt32(&i, 1)
}

func sub() {
	atomic.AddInt32(&i, -1)
}

func load() {
	atomic.LoadInt32(&i) //读的时候可以用
}

func store() {
	atomic.StoreInt32(&i, 100) //写的时候用
}

func main() {
	for j := 0; j < 100; j++ {
		go add()
		go sub()
	}
	runtime.Gosched()
	fmt.Println(i)
}
