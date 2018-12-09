package main

import (
	"fmt"
)

func foo() {
	fmt.Println("Foo function ran")
}

func main() {
	foo()

	func() {
		fmt.Println("Anonymous function ran")
	}()

	func(x int) {
		fmt.Println("The value received by anonymous function is: ", x)
	}(999)
}
