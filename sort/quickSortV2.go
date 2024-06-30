package main 

import(
	"fmt"
)

// 定义快排结构题
type quickSorter struct {
	// 记录递归层级，便于分析解释为什么快排不稳定，可能出现O(N^2)的时间复杂度
	// 会用 升序数组，使用最左侧元素作为 基准值做升序排序， 浮现O(N^2)时间复杂度
	level int
}

// 快速排序方法
// 思路分治
// 1. 选取基准值 遍历保证   左边数据<=基准值<=右边数据
// 2. 接着拆分成为两个数组, 重复1 操作 最终递归结束
func (qs *quickSorter) quickSort(arr []int,  left , right , level int) {

	// 递归出口
	if left >= right {
		return 
	}

	
	// 记录当前递归拆分的最大深度 
	qs.level = qs.max(qs.level, level)
	// 尽可能选取中间值 作为基准值
	mid := (left + right) >> 1
	midIndex := qs.getMidIndex(arr, left, right, mid)

	// 选择最左边下标作为基准值
	// midIndex := left
	
	l := left
	r := right

	for l < r {
		for l < r && arr[l] < arr[midIndex] {
			l ++
		}
		for l < r && arr[r] >= arr[midIndex] {
			r --
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[midIndex], arr[l] = arr[l], arr[midIndex]
	
	// 分治重复上述操作
	qs.quickSort(arr, left, l - 1, level + 1)
	qs.quickSort(arr, l + 1, right, level + 1)

	return

}

// 选取中位数做基准值
func (qs *quickSorter) getMidIndex(arr []int, left, right, mid int) int {

	if arr[left] >= arr[right] && arr[left] >= arr[mid]  {
		return left
	}

	if arr[right] >= arr[left] && arr[right] >= arr[mid]  {
		return right
	}

	return mid
}



func (qs *quickSorter) max(a, b int) int{
	if a >= b  {
		return a
	} else {
		return b
	}
}	

func main () {

	qs := quickSorter{0}
	
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("before sort: ", arr)
	qs.quickSort( arr, 0, len(arr) -1, 0)
	fmt.Println("after sort: ", arr)
	fmt.Println("max level: ", qs.level)
}