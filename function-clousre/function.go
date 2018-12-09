package main

import (
	"fmt"
)

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	var x int
	fmt.Println(x)
	x++
	// code block
	{
		y := 990
		fmt.Println(y)
	}
	fmt.Println(x)

	a := increment()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
}
