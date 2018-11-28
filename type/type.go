package main

import (
	"fmt"
)

var x = 10
var name = "Vaibhav Singh"

// can also declare like this
var z string = "random string"
var y = `some string, "random string"`

func main() {
	fmt.Println(x)
	fmt.Printf("Type: %T\n", x)
	fmt.Println(name)
	fmt.Printf("Type:  %T\n", name)
	fmt.Println(z)
	fmt.Printf("Type: %T\n", z)
	fmt.Println(y)
	fmt.Printf("Type: %T\n", y)
	// will give error as declared string but assigning int
	// Go is a static language not dynamic programming language

	// name = 100
	// fmt.Println(name)
}
