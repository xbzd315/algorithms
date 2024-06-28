package main 

import(
	"fmt"
)

func main () {
	
	var a int32 = 1000
	
	var b int8 = int8(a)
	
	var c int64 = int64(a)

	fmt.Printf("a type :%T a value %v \n b type :%T b value:%v \n c type: %T c value: %v\n", a, a, b, b, c, c)
}