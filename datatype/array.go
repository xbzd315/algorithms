package main

import (
	"fmt"
)

func main() {

	// var arr1 = [6]float64 {1, 2, 3, 4, 5 ,6}
	var arr1 = [...]float64 {1, 2, 3, 4, 5 ,6}

	fmt.Println(" arr1 :", arr1)

	var total float64 = 0
	// 遍历数组
	for  _ , val  := range arr1 {
		val += 1
		total += val
	}

	// 数组类型在内存中是顺序存储的，数组的地址就是数组第一个元素的地址，后续元素的地址都是一次按照数据类型地址累加的
	fmt.Printf(" arr1 地址 %p:\n arr1[0] 地址 %p\n arr1[1] 地址 %p\n", &arr1, &arr1[0], &arr1[1])
	fmt.Println(" arr1 :", arr1)
	fmt.Println(" total :", total)

	for i := 0; i < len(arr1) ; i ++ {
		arr1[i] += 1
	}

	fmt.Println(" arr1 :", arr1)

}	