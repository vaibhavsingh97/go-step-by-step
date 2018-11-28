package main

import (
	"fmt"
)

// Can declare variable outside a function
var myName = "Vaibhav Singh"

// Delcare VARIABLE z with identifier int, assigns 0 value automatically
// as it's default value
var z int
var y float32

func main() {
	// Can only declare inside a function
	x := 100
	fmt.Println(x)
	fmt.Println(myName)
	foo()
	fmt.Println(z)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Printing my name again:", myName)
}
