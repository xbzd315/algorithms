package main
import(
	"fmt"
	"strconv"
)
func main() {
	var a int = 123
	var b float64 = 3.1415
	var c bool = true
	var d byte = '?'
	var e string
	e = fmt.Sprintf("%d", a)
	fmt.Printf("e type %T e value %v \n", e, e)

	e = fmt.Sprintf("%f", b)
	fmt.Printf("e type %T e value %v \n", e, e)

	e = fmt.Sprintf("%t", c)
	fmt.Printf("e type %T e value %v \n", e, e)

	e = fmt.Sprintf("%c", d)
	fmt.Printf("e type %T e value %v \n", e, e)

	v := uint64(42)
	// 数字转10进制字符串
	s1 := strconv.FormatUint(v, 10)
	fmt.Printf("e type %T e value %s \n", s1, s1)

	s2 := "-354634382"
	if s, err := strconv.ParseInt(s2, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}