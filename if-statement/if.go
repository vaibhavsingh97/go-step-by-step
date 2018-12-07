package main

import (
	"fmt"
)

func main() {
	x := 40
	if x == 40 {
		fmt.Println("001")
	}
	if x != 40 {
		fmt.Println("002")
	}

	if x <= 50 {
		fmt.Println("003")
	}
	if x >= 0 {
		fmt.Println("004")
	}

}
