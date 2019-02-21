package main

import (
	"errors"
	"fmt"
)

/**
异常错误处理相关demo
*/

func err(a int) (int, error) {
	if a == 42 {
		return -1, errors.New("这里有异常")
	}
	return a, nil
}

/**
自定义错误类型
*/
type myError struct {
	arg  int
	prob string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func err2(arg int) (int, error) {
	if arg == 42 {
		return -1, &myError{arg, "can not work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := err(i); e != nil {
			fmt.Println("err", e)
		} else {
			fmt.Println(r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := err2(i); e != nil {
			fmt.Println("err2", e)
		} else {
			fmt.Println(r)
		}
	}
	/**
	如果需要在程序中使用一个自定义的错误类型中的数据，你需要通过类型断言来得到这个错误类型的实例
	*/
	// _,e := err2(42)
	//if ae,ok := e.(*myError); ok {
	//	fmt.Println(ae.arg)
	//	fmt.Println(ae.prob)
	//}
}
