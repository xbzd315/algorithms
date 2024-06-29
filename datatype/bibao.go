package main
import(
	"fmt"
)
func add() func (int) int {

	var n = 10

	return func (m int) int {
		n += m
		return n 
	}
}

func main () {

	func1 := add() 
	fmt.Println(func1(1))
	fmt.Println(func1(1))
	fmt.Println(func1(1))

}