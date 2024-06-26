package main
import (
    "fmt"
    "unsafe"
)
func main() {
    var b bool
	b = true
    fmt.Printf("Size of bool: %d byte(s)\n", unsafe.Sizeof(b))
    fmt.Println("" ,b)
}
