package main

import (
	"fmt"
)

func main() {
	// x := type{values} // composite literal
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[:1])
	fmt.Println(x[:])
	fmt.Println(x[1:3])
}

// SLICE allows you to group together Values of same type
