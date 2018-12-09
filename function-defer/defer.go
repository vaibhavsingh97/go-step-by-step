package main

import (
	"fmt"
)

func main() {
	// We use defer to close or deallocate resources usually. It is guaranteed
	//to run in every case. In case of any return or panic. You don't have this
	// guarantee when you just place the function at the end.
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
