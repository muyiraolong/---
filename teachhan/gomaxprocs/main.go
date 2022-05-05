package main

import (
	"fmt"
	"runtime"
	"time"
)

func a(u int) {
	for i := 0; i < 10; i++ {
		fmt.Println(u, "\t", i)
	}
}

func b(s int) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "\t", i)
	}
}

func main() {
	// fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(12)
	for i := 0; i < 20; i++ {
		go a(100 + i)
		go b(10 + i)
	}
	time.Sleep(time.Second)
}
