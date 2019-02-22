package main

import (
	"fmt"
	"time"
)

/**
工作池
*/

func pool_worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	/**
	执行时通道还没有数据
	*/
	for w := 1; w <= 3; w++ {
		go pool_worker(w, jobs, results)
	}

	/**
	此时往通道里面添加数据
	*/
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	/**
	收集任务的返回值
	*/
	for a := 1; a <= 9; a++ {
		<-results
	}
}
