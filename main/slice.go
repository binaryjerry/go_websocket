package main

import "fmt"

/**
Slice是go中一个关键的数据类型，是一个比数组更加强大的序列接口

跟数组不一样的是，slice的类型仅仅由它所包含的元素决定。
*/
func main() {

	//创建一个长度非0的空slice，需要使用内建的方法make
	s := make([]string, 3)
	fmt.Println("emp:", s)
	//赋值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	//取值
	fmt.Println("get:", s[1])
	fmt.Println("length", len(s), "这长度")
	//追加
	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(len(s))
	//Slice也可以被copy。
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	//真正意义上的切片操作
	l := s[2:5]
	fmt.Println("slice l", l)

	l2 := s[:5]
	fmt.Println(l2)

	l3 := s[2:]
	fmt.Println(l3)

	//直接声明有值得Slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	//同时Slice也可以构成多维数据结构,内部的Slice长度可以不同。这和多维数组不同
	TwoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		TwoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			TwoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", TwoD)
}
