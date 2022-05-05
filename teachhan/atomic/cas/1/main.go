// atomic底层本身就是用的cas
package main

import (
	"fmt"
	"sync/atomic"
)

var i int32 = 100

func main() {
	b := atomic.CompareAndSwapInt32(&i, 100, 200) //交换一个值，比较看是否交换成功
	fmt.Println(b, i)
}
