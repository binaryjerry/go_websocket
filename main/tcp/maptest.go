package tcp

import (
	"fmt"
	"go/types"
	"reflect"
)

//定义map
type Param map[string]interface{}

//嵌入Param到struct
type Show struct {
	Param
}

var (
	array  = [2]int{2, 1}
	sliceS types.Slice
)

func main() {
	//初始化结构体
	s := new(Show)
	//初始化结构体内嵌入的map
	s.Param = Param{}
	//给结构体初始化值
	s.Param["RMB"] = 10000
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(Param{}))
	fmt.Println(reflect.TypeOf(array))

}
