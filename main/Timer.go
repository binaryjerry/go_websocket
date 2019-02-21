package main

import (
	"fmt"
	"time"
)

/**
定时器
*/
func main() {

	timer1 := time.NewTimer(time.Second * 2)

	//直到这个定时器的通道C明确的发送了定时器失效的值之前，将一直阻塞
	<-timer1.C

	fmt.Sprintf("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stoppped")
	}
}
