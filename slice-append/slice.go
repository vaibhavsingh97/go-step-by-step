package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	x = append(x, 33, 44, 55, 66)
	fmt.Println(x)
	y := []int{99, 88, 77, 66}
	// if  "..." before some element than it will take unlimited number of elements
	// if "..." after the element than it will expand all the elements
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...) // delete 3,5 from slice
	fmt.Println(x)
}
