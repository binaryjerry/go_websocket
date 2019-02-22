package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

/**
原子计数器
*/
func main() {

	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			//运行其它协程执行
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops", opsFinal)
}
