package main

import (
	"fmt"
)

func main() {

	// map 使用 
	var a map[string] string
	// 使用map前 需要make分配空间
	a = make(map[string] string , 10)

	a["123"] = "测试"
	a["124"] = "测试1"
	var s string = "123"
	fmt.Println(" key => value :" + s + a[s])
}