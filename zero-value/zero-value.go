package main

import (
	"fmt"
)

// if we don't assign a value to variable, go automatically assign zero value
var y string
var z int
var x bool
var a float64

func main() {
	fmt.Println("printing value of y", y, "end")
	fmt.Printf("Type: %T\n", y)
	fmt.Println("printing value of z", z, "end")
	fmt.Printf("Type: %T\n", z)
	fmt.Println("printing value of x", x, "end")
	fmt.Printf("Type: %T\n", x)
	fmt.Println("printing value of a", a, "end")
	fmt.Printf("Type: %T\n", a)
}
