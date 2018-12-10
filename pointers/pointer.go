package main

import (
	"fmt"
)

func main() {
	x := 23
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Printf("%T\n", &x)

	y := &x
	fmt.Println(y)
	fmt.Println(*y)

	// & -> Gives the address
	// * -> de reference the address and gives value stored of the address

	*y = 24
	// x will be 24
	// As x and y sharing the same house because of pointers
	fmt.Println(x)
}
