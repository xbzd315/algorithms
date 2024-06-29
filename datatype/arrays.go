package main

import (
	"fmt"
)

func main() {

	var arrs [10][10]int64

	// fmt.Println(" arrs :", arrs)

	for index, vals := range arrs {

		fmt.Printf("index = %v  vals : %v 地址 %p \n", index, vals, &arrs[index])
		for _, val := range vals  {
			val += 1
			// fmt.Println(" val :", val)
		}
		fmt.Printf(" arrs[%v][9]  地址: %p \n", index, &arrs[index][9])	
		
		if index >= 2 {
			break;
		}
	}
}	