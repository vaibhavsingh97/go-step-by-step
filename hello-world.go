package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello World! I am learning Go :)")
	foo()
	fmt.Println("Loop ahead")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	exit()
}

func foo() {
	fmt.Println("I'm in foo")
}

func exit() {
	fmt.Println("Exiting ...")
}
