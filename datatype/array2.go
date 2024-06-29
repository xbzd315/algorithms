package main

import (
	"fmt"
)


func update (a [6]float64) {

	a[0] = 10

	return
}


func updatebyPointer (a *[6]float64) {
	fmt.Printf("a value %p *a value :%v a 地址 : %p \n", a, *a, &a)
	(*a)[0] = 10
	// a[0] = 10

	return
}

func main() {

	var arr1 = [...]float64 {1, 2, 3, 4, 5 ,6}
	var arr1 = [6]float64 {1, 2, 3, 4, 5 ,6}
	
	fmt.Println(" arr1 :", arr1)
	
	update(arr1)
	fmt.Println("after update arr1 :", arr1)

	fmt.Printf("arr1 value :%v arr1 地址 : %p \n", arr1, &arr1)
	updatebyPointer(&arr1)
	fmt.Println("after updatebyPointer arr1 :", arr1)


}	