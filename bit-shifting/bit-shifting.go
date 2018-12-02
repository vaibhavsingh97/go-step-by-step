package main

import (
	"fmt"
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)

	// print even number using bit shifting
	for i := 0; i < 10; i++ {
		fmt.Println(i << 1)
	}

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%T\t\t%b\n", kb, kb)
	fmt.Printf("%T\t\t%b\n", mb, mb)
	fmt.Printf("%T\t\t%b\n", gb, gb)

	// Another way

	const (
		_   = iota
		kb1 = 1 << (iota * 10)
		mb1 = 1 << (iota * 10)
		gb1 = 1 << (iota * 10)
	)

	fmt.Printf("%T\t\t%b\n", kb1, kb1)
	fmt.Printf("%T\t\t%b\n", mb1, mb1)
	fmt.Printf("%T\t\t%b\n", gb1, gb1)
}
