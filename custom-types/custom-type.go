package main

import (
	"fmt"
)

var x int

type apple int

var y apple

func main() {
	x = 100
	y = 10001

	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)

	fmt.Println(y)
	fmt.Printf("Type: %T\n", y)

	// This will give error as TYPE apple cannot be assigned to TYPE int
	// for assignments both have to be have same types
	// x = y --> Error
	// conversion
	x = int(y)
	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)
}
