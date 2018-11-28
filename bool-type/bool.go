package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	foo()
}

func foo() {
	a := 10
	b := 100
	fmt.Println(a == b)
}
