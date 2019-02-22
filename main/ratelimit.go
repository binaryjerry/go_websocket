package main

import (
	"fmt"
	"time"
)

/**
速率限制
*/
/**
速率，去控制服务资源利用和质量的途径
*/
func main() {

	//定义一个int类型的，容量为5的通道
	requests := make(chan int, 5)

	//往通道输入1~5数字
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	//关闭通道
	close(requests)

	//速率限制任务中的管理器
	limiter := time.Tick(time.Millisecond * 200)

	//通过每次请求前阻塞limiter通道的一个接收，我们限制自己每200ms执行一次请求
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	//每200ms添加一个新的值到burstyLimiter通道
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	//模拟5个请求
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
