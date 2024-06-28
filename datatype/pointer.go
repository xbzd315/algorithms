package main
import(
	"fmt"
)
func main() {
	var a int = 100
	// var b int 
	// var c string 
	// 定义指针类型
	var d *int 
	fmt.Printf(" a value %v \n a dizhi %v \n", a, &a)
	d = &a
	*d = 30
	fmt.Printf(" d value %v \n d dizhi %v  \n d -> value %v  \n", d, &*&*&d, *d)
}

// e := []int{1,2,3,5}
// var f *[]int 
// f = &e
// fmt.Printf(" f value %v \n f dizhi %v  \n f -> value %v  \n", f, &*&*&f, *f)