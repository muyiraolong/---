package main

import (
	"log"
	"time"
)

func main() {
	// time1 := time.NewTimer(time.Second)
	// fmt.Println(time.Now())
	// t1 := <-time1.C //这是一个阻塞通道，直到时间停止
	// fmt.Println(t1)

	// time2 := time.NewTimer(time.Second)
	// <-time2.C //=time.Sleep(time.Second)

	// fmt.Println(time.Now())
	// <-time.After(time.Second)
	// fmt.Println(time.Now())

	// time3 := time.NewTimer(time.Second)
	// go func() {
	// 	<-time3.C
	// 	fmt.Println("没阻塞")
	// }()
	// s := time3.Stop() //直接让time3停掉，并停掉协程
	// if s {
	// 	println(s)
	// }
	// time.Sleep(time.Second)

	log.Println("before")
	timer4 := time.NewTimer(time.Second * 5) //原来设置5s
	timer4.Reset(time.Second * 1)            //重新设置时间,即修改NewTimer的时间
	<-timer4.C
	log.Println("after")
}
