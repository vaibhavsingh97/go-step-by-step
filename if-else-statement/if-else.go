package main

import (
	"fmt"
)

func main() {
	x := 40
	if x == 40 {
		fmt.Println("x value is 40")
	} else if x == 41 {
		fmt.Println("x value is not 41")
	} else {
		fmt.Println("x value was not 40 or 41")
	}
}
