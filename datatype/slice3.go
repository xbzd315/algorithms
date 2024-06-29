package main

import "fmt"

func main() {
 	// 方式一
    // 使用 make 创建切片
	// 创建一个长度为 5，容量为 5 的整型切片
    slice1 := make([]int, 5)
	// 创建一个长度为 3，容量为 5 的整型切片
    slice2 := make([]int, 3, 5)
}