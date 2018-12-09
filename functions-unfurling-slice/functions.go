package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []int{}
	// code on L9 gives error as "int" and "slice of int" is different
	// we have to unload the slice of int
	// "..." aka variadic means 0 or more
	// sum_of_x := sum(x)
	sumOfX := sum(x...)
	// len and cap will be 0
	sumOfY := sum(y...)
	fmt.Println(sumOfX)
	// sum will be 0
	fmt.Println(sumOfY)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
