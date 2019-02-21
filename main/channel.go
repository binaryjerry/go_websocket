package main

import (
	"fmt"
	"time"
)

/**
通道相关demo
*/

/**
通道是连接多个协程的管道，你可以从一个Go协程将值发送到通道，然后在别的Go协程中接收
*/

func main() {

	//创建一个通道
	messages := make(chan string)

	//往通道里面传值
	go func() { messages <- "ping" }()

	//从通道里面取值
	msg := <-messages
	fmt.Println(msg)

	/**
	通道缓冲
	*/
	messages2 := make(chan string, 2)

	messages2 <- "buffered"
	messages2 <- "channel"

	fmt.Println(<-messages2)
	fmt.Println(<-messages2)

	//通道同步
	done := make(chan bool, 1)
	go worker(done)
	//程序将在接收到通道中worker发出的通知前一直阻塞
	<-done

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	/**
	通道选择器
	*/
	/**
	通道选择器让你同时等待多个通道操作，Go协程和通道以及选择器的结合是Go的强大特性
	*/
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	//使用select关键字来同时等待这两个值，并且打印
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	/**
	超时处理
	*/
	c3 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c3 <- "result 1"
	}()

	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	/**
	通道的遍历
	*/
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

/**
通道同步
*/
/**
我们可以使用通道来同步Go协程间的执行的状态，这里是一个使用阻塞的接收方式来等待一个GO协程的运行结束
*/
func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	//发送一个值来通知我们已经完成这个函数了
	done <- true
}

/**
通道方向
*/
/**
当使用通道作为函数参数时，你可以指定这个通道是不是只用来发送或者接收值，这个特性提升了程序的类型安全性
*/

//只发送
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//运行pings发送，运行pongs接收
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
