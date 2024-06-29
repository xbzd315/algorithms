package main

import "fmt"

func main() {

    // 从数组创建切片
    array := [5]int{1, 2, 3, 4, 5}
    slice1 := array[1:4] // 包含 array[1] 到 array[3] 的元素 [1,4) 左闭右开
    slice2 := array[1:3] // 包含 array[1] 到 array[2] 的元素 [1,3) 左闭右开

	// 只改了array， 但是其他的两个切边是引用的array所以 都会跟着改变。
	array[1] = 100

	fmt.Printf("array[1] 的地址是 %p 值为:%v \n", &array[1], array[1])
	fmt.Printf("slice1[0] 的地址是 %p 值为:%v \n", &slice1[0], slice1[0])
	fmt.Printf("slice2[0] 的地址是  %p 值为:%v \n", &slice2[0], slice2[0])

}