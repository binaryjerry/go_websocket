package main

import "fmt"

/**
range遍历，可以迭代各种数据
*/

func main() {

	//1.使用range来统计一个slice的元素个数
	nums := []int{1, 2, 3, 4}
	sum := 0
	//忽略索引的做法
	for _, num := range nums {
		sum += num
	}

	//带索引的做法
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	//range在map中的使用
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	//range在字符串中迭代unicode编码，第一个返回的是rune的起始字节位置，然后第二个是rune自己
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	fmt.Println("sum:", sum)
}
