package util

import(
	"fmt"
)

var A int = 1

func init() {
	fmt.Println(" init func A value", A)
	A = 2
	fmt.Println(" init func A value", A)
}