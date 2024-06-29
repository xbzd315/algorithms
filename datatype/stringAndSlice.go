package main

import (
	"fmt"
)

func main() {
	// string 底层是一个byte数组， 因此可以进行切片处理
	fmt.Printf("-----string 底层是一个byte数组， 因此可以进行切片处理-----")
	str := "abceetst:12312!啊"
	// 使用切片 获取下标9及后面的元素
	slice := str[10:]
	var a byte 
	fmt.Printf("a type: %T \n", a)
	fmt.Printf("slice[0] type: %T \n", slice[0])
	fmt.Println("slice= ", slice)
	// slice[0] += 
	//  string 是不可变的，str[0] 去赋值是不可以的
	// fmt.Println("-----string 是不可变的，str[0] 去赋值是不可以的-----")
	// str[0] = "c" //./stringAndSlice.go:22:2: cannot assign to str[0] (value of type byte)

	fmt.Println("-----string 修改示例-----")
	slice1 := []byte(str)
	slice1[0] = 98 // b = 98  
	str = string(slice1)
	fmt.Println("str=", str)

	slice2 := []rune(str)
	slice3 := str[:]
	slice2[0] = '啊' // b = 98  
	str = string(slice2)
	fmt.Println("str=", str)
	fmt.Println("slice2=", slice2)
	fmt.Println("slice3=", slice3)
	// fmt.Printf("str[0] 地址= %p \n",  &str[0])
	// fmt.Printf("slice3[0] 地址= %p \n",  &slice3[0])
	fmt.Printf("slice2[0] 地址= %p \n",  &slice2[0])
}