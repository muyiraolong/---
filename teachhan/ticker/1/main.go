// 可周期执行的timer
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker1 := time.NewTicker(time.Second)
	// counter := 1
	// for range ticker1.C {
	// 	log.Println("1")
	// 	counter++
	// 	if counter > 5 {
	// 		ticker1.Stop()
	// 		break //需要手动break，原理是，ticker停止周期执行了，但for range还在询问，如果不break，会panic
	// 	}
	// }

	// chanInt := make(chan int)
	// 	go func() {
	// 		sum := 0
	// 		for v := range chanInt {
	// 			sum += v
	// 			fmt.Println(sum)
	// 			if sum > 10 {
	// 				break
	// 			}
	// 		}
	// 	}()
	// 	for range ticker1.C {
	// 		select {
	// 		case chanInt <- 1:
	// 			fmt.Println("send: 1")
	// 		case chanInt <- 1:
	// 			fmt.Println("send: 2")
	// 		default:
	// 			fmt.Println("send 3")
	// 		}
	// 	}

	// }

	// counter := 1
	// for range ticker1.C {
	// 	log.Println("1")
	// 	counter++
	// 	if counter > 5 {
	// 		ticker1.Stop()
	// 		break //需要手动break，原理是，ticker停止周期执行了，但for range还在询问，如果不break，会panic
	// 	}
	// }

	chanInt := make(chan int)
	go func() {
		for range ticker1.C {
			select {
			case chanInt <- 1:
				fmt.Println("send: 1")
			case chanInt <- 1:
				fmt.Println("send: 2")
			default:
				fmt.Println("send 3")
			}
		}
	}()
	sum := 0
	for v := range chanInt {
		sum += v
		fmt.Println(sum)
		if sum > 10 {
			break
		}
	}

}

// 总的来说，需要break来打断继续执行time.Ticker，
// 如果这里把time.ticker放在主协程，下面的放在go func中，主协程就会继续访问。这里就相当于主协程跑完了，就停掉了go func协程
