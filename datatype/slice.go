package main

import "fmt"

func main() {
    // 使用 make 创建切片
	// 创建一个长度为 5，容量为 5 的整型切片
    slice1 := make([]int, 5)
	// 创建一个长度为 3，容量为 5 的整型切片
    slice2 := make([]int, 3, 5)
    
    // 使用字面量创建切片
    slice3 := []int{1, 2, 3, 4, 5}
    
    // 从数组创建切片
    array := [5]int{1, 2, 3, 4, 5}
    slice4 := array[1:4] // 包含 array[1] 到 array[3] 的元素
    
    // 切片的切片
    slice5 := []int{1, 2, 3, 4, 5}
    slice6 := slice5[1:4]// 包含 slice5[1] 到 slice5[3] 的元素
    
    // 零值切片
    var slice7 []int// slice7 是一个 nil 切片
    
    // 使用 append 创建切片
    var slice8 []int
    slice8 = append(slice8, 1, 2, 3)
    
    // 打印所有切片
    fmt.Println("slice1:", slice1)
    fmt.Println("slice2:", slice2)
    fmt.Println("slice3:", slice3)
    fmt.Println("slice4:", slice4)
    fmt.Println("slice5:", slice5)
    fmt.Println("slice6:", slice6)
    fmt.Println("slice7:", slice7)
    fmt.Println("slice8:", slice8)
}
