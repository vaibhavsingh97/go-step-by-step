package main

import (
	"fmt"
)

var x = 10

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", x)
	fmt.Printf("%#x\n", x)

	// String printing
	s := fmt.Sprintf("%#x\t%b\n", x, x)
	fmt.Println(s)

	// Print value
	p := fmt.Sprintf("%v\n", x)
	fmt.Println(p)

}
