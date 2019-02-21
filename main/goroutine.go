package main

import "fmt"

/**
协程相关demo
*/

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Sprintf(from, ":", i)
	}
}

func main() {
	//一般方法
	f("direct")

	////使用协程
	//go f("goroutine")
	//
	////为匿名函数启动一个Go协程
	//go func(msg string) {fmt.Println(msg)}("going")

	//var input string

	//fmt.Scanln(&input)
	//fmt.Println("done")

}
