package main

import (
    "fmt"
    _"math/rand"
    _"time"
    qs "github.com/xbzd/algorithms/quickSort"
)

// // 定义一个结构体 quickSort
// type quickSort struct{
// 	level int
// }
// // 定义一个结构体 quickSortV2
// type quickSortV2 struct{
// 	level int
// }
// // 定义一个结构体 mergeSort
// type mergeSort struct {
//     level int
// }
// // partition 方法，使用随机基准数进行哨兵划分
// func (q *quickSort) partition(nums []int, left, right int) int {
//     // 随机选择一个基准数
//     // rand.Seed(time.Now().UnixNano())
//     // pivotIndex := left + rand.Intn(right-left+1)
// 	pivotIndex := left
//     // 将基准数交换到数组的最左端
//     nums[left], nums[pivotIndex] = nums[pivotIndex], nums[left]
    
//     i, j := left, right
//     for i < j {
//         // 从右向左找首个小于基准数的元素
//         for i < j && nums[j] >= nums[left] {
//             j--
//         }
//         // 从左向右找首个大于基准数的元素
//         for i < j && nums[i] <= nums[left] {
//             i++
//         }
//         // 交换找到的两个元素
//         nums[i], nums[j] = nums[j], nums[i]
//     }
//     // 将基准数交换至两子数组的分界线
//     nums[i], nums[left] = nums[left], nums[i]
//     return i // 返回基准数的索引
// }

// // quickSort 方法，对数组进行快速排序
// func (q *quickSort) quickSortV1(nums []int, left, right int) {
//     // 子数组长度为 1 时终止递归
//     if left >= right {
//         return
//     }
// 	fmt.Println("%c", q.level)
// 	q.level += 1
//     // 哨兵划分
//     pivot := q.partition(nums, left, right)
//     // 递归左子数组、右子数组
//     q.quickSortV1(nums, left, pivot-1)
//     q.quickSortV1(nums, pivot+1, right)
// }


// /* 选取三个候选元素的中位数 */
// func (q *quickSortV2) medianThree(nums []int, left, mid, right int) int {
//     l, m, r := nums[left], nums[mid], nums[right]
//     if (l <= m && m <= r) || (r <= m && m <= l) {
//         return mid // m 在 l 和 r 之间
//     }
//     if (m <= l && l <= r) || (r <= l && l <= m) {
//         return left // l 在 m 和 r 之间
//     }
//     return right
// }

// /* 哨兵划分（三数取中值）*/
// func (q *quickSortV2) partition(nums []int, left, right int) int {
//     // 以 nums[left] 为基准数
//     med := q.medianThree(nums, left, (left+right)/2, right)
//     // 将中位数交换至数组最左端
//     nums[left], nums[med] = nums[med], nums[left]
//     // 以 nums[left] 为基准数
//     i, j := left, right
//     for i < j {
//         for i < j && nums[j] >= nums[left] {
//             j-- //从右向左找首个小于基准数的元素
//         }
//         for i < j && nums[i] <= nums[left] {
//             i++ //从左向右找首个大于基准数的元素
//         }
//         //元素交换
//         nums[i], nums[j] = nums[j], nums[i]
//     }
//     //将基准数交换至两子数组的分界线
//     nums[i], nums[left] = nums[left], nums[i]
//     return i //返回基准数的索引
// }

// // quickSort 方法，对数组进行快速排序
// func (q *quickSortV2) quickSortV2(nums []int, left, right int) {
//     // 子数组长度为 1 时终止递归
//     if left >= right {
//         return
//     }
// 	fmt.Println("level:", q.level)
// 	q.level += 1
//     // 哨兵划分
//     pivot := q.partition(nums, left, right)
//     // 递归左子数组、右子数组
//     q.quickSortV2(nums, left, pivot-1)
//     q.quickSortV2(nums, pivot+1, right)
// }


// 主函数
func main() {
    // nums := []int{3, 6, 8, 10, 1, 2, 1}
    nums := []int{3, 1, 0, 4, 5, 2, 9, 12}
    
    // 快速排序版本1
    // sorter := quickSort{1}
	// sorter.level = 1
    // sorter.quickSortV1(nums, 0, len(nums)-1)

    // 快速排序版本2
	sorter := sq.quickSortV2{}
	sorter.level = 1
    sorter.quickSortV2(nums, 0, len(nums)-1)
    fmt.Println(nums)
}
