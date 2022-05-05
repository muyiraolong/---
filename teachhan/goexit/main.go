package main

import (
	"fmt"
	"runtime"
	"time"
)

func show(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "\t", i)
		if i > 6 {
			runtime.Goexit()
		}
	}
}
func main() {
	go show("JAVA")
	time.Sleep(time.Second)

}
