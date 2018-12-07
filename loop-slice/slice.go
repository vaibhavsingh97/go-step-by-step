package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[2])
	//for loop for slice
	for index, value := range x {
		fmt.Println(index, value)
	}
	fmt.Println("===========Another method===========")
	// another method to loop a slice
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
