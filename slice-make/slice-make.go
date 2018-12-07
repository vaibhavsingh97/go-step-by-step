package main

import (
	"fmt"
)

func main() {
	// It helps prevent recreating arrays at runtime
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[0] = 999
	x[9] = 1000
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 87987987)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
