package main

import (
	"fmt"
)

func foo() string {
	return "Hello World"
}

func bar() func() int {
	return func() int {
		return 9999
	}
}

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := bar()
	fmt.Printf("%T\n", s2)
	fmt.Println(s2)
	fmt.Println(s2())
	fmt.Println(bar()())
}
