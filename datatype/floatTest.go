package main

import (
    "fmt"
)

func main() {
    // 定义一个浮点数
    var num float64 = 123.0000905

    // 将浮点数分别转换为 float32 和 float64
    var num32 float32 = float32(num)
    var num64 float64 = num

    // 打印两个浮点数的值
    fmt.Printf("Original number: %.7f\n", num)
    fmt.Printf("float32: %.7f\n", num32)
    fmt.Printf("float64: %.7f\n", num64)
}