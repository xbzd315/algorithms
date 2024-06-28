package main

import (
    "fmt"
)

// 字符串类型使用
func main() {
	// ascii 值
	var c1 byte = '0'
	// c1 += 1
	// var c2 byte = '杨'
	var c2 int16 = '杨'
	c2 += 1;
	// var c2 int = "杨某"
	
	var c3 string = "杨某"

	// ascii 值
	fmt.Println("c1 =", c1)
	// %c按照字符输出
	fmt.Printf("c1 =%c\n", c1)
	// %d 按照数字输出
	fmt.Printf("c2 =%c\n c2= %d\n", c2, c2)

	fmt.Printf("c3 =%c\n c3= %d\n", c3, c3)
	// 
}