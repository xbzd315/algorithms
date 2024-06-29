package main

import(
	"fmt"
)

func main() {

	var a int = 10;
	// 执行到defer时， 会将需要执行的语句压入到defer独立栈的中，暂时不执行。
	// 当函数执行完的时候，从defer栈中，按照栈的策略出栈执行。
	// 按道理应该输出的是 10，而不是
	defer  fmt.Println("befroe add a value-1", a)
	defer  fmt.Println("befroe add a value-2", a)
	
	a += 1
	fmt.Println("after add a value", a)

}