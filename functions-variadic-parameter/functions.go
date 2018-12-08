package main

import (
	"fmt"
)

func main() {
	x := foo(1, 2, 3)
	fmt.Println("The total sum is: ", x)
}

// this function will return "slice" of int
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
