package main

import (
	"fmt"
)

// 定义归并排序 结构体
type mergeSort struct {
	// 用于计算递归的最大深度
	level int
}
// 归并排序方法
func (ms *mergeSort) mergeSort(arr, temp []int, left ,right, level int) {

	// 递归出口，拆分到最小粒度从这里开始回溯
	if left >= right {
		return 
	}
	mid := (left + right) >> 1

	// 记录当前递归拆分的最大深度 
	ms.level = ms.max(ms.level, level)
	// 拆分
	ms.mergeSort(arr, temp, left, mid, level+1)
	ms.mergeSort(arr, temp, mid+1, right, level+1)

	l := left
	l1 := mid + 1
	// 回溯合并的过程
	for i := left ; i <= right ; i++ {
		if l > mid  {
			temp[i] = arr[l1]
			l1 += 1
			continue
		}
		if l1 > right  {
			temp[i] = arr[l]
			l += 1
			continue
		}
		// 按照顺序赋值 给临时切片
		if arr[l] <= arr[l1]  {
			temp[i] = arr[l]
			l += 1
		} else {

			temp[i] = arr[l1]
			l1 += 1
		}
	}

	// 临时切片按照顺序复制给 排序切片 
	for i := left ; i <= right ; i++ {
		arr[i] = temp[i]
	}

	return 
}

// 求最大值
func (ms *mergeSort) max(a, b int) int{
	if a >= b  {
		return a
	} else {
		return b
	}
}	

func main () {

	//  定义需要排序的 数组
	// var arr []int = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr  := []int{1, 2, 3, 4, 5, 6, 8, 7, 9, 10}
	fmt.Println("before sort: ", arr)
	var temp []int = make([]int , len(arr))

	// 定义结构体类型
	ms := mergeSort{0}
	ms.mergeSort(arr, temp, 0, len(arr)-1, 0)
	fmt.Println("after sort: ", arr)
	fmt.Println("max level: ", ms.level)
}