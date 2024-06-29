package main

import(
	"fmt"
)


var (
	// 定义 全局匿名函数
	Fun = func (a ,b int) int {
		return a - b
	}
)

func main () {


	//  匿名函数方式 直接返回结果
	ret := func (a ,b int) int {
		return a + b
	}(1, 2)
	fmt.Println("ret :", ret)

	// 匿名函数赋值给变量
	func1 := func (a ,b int) int {
		return a + b
	}
	ret1 := func1(2, 3)
	fmt.Println("ret1 :", ret1)

	// 全局匿名函数使用
	ret2 := Fun(2, 3)
	fmt.Println("ret2 :", ret2)
}