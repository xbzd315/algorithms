package main

import (
    "fmt"
)

// 字符串类型使用
func main() {


	
	var c3 string = "abcde杨"

	// golang 用的 utf-8 所以中午字符占了3个字节， golang遍历是一个字节一个字节遍历
	fmt.Println("c3 len =", len(c3))

	// 将含有中文字符的字符串转换成 rune的切片
	c4 := []rune(c3)
	for i := 0; i < len(c4) ; i++ {
		fmt.Printf("val =%c\n", c4[i])
	}


	for index, val := range c3 {
		fmt.Printf("index val = %v, %c", index, val)
	}
}