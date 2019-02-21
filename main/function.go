package main

import "fmt"

//单返回值函数
func plus(a int, b int) int {
	return a + b
}

//多返回值函数
func mutilPlus(a int, b int) (int, int) {
	return a + b, a - b
}

//变参函数
//这个函数表示使用任意数目的int作为参数
func sum(nums ...int) {
	fmt.Println(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(mutilPlus(1, 2))
	sum(1, 2, 3, 4, 5, 6)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	sum(nums...)
}
