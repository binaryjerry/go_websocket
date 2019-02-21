package main

import "fmt"

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
}
