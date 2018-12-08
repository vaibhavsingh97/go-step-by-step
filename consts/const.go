package main

import (
	"fmt"
)

type foo int

var y foo

const bar = 42

func main() {
	y = bar
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", bar)
	fmt.Println(bar)
}
