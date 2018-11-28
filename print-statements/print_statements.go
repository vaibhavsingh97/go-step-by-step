package main

import (
	"fmt"
)

func main() {
	res, e := fmt.Println("Hey, I am new to go")
	fmt.Println(res)
	fmt.Println(e)

	foo()
	foo1()
}

func foo() {
	res, _ := fmt.Println("I am in foo")
	fmt.Println(res)
}

func foo1() {
	// This will give error as we had declared "e" but not used
	res, e := fmt.Println("I am in Foo 1")
	fmt.Println(res)
}
