package main

import (
	"fmt"
	"runtime"
)

func show(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, "\t", i, "\t")
	}
}

func main() {
	go show("JAVA")
	for i := 0; i < 2; i++ {
		runtime.Gosched() //让子协程进行
		fmt.Println(i)
	}
}
