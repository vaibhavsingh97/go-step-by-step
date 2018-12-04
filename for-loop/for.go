package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// Example of break and continue
	z := 1
	for {
		z++
		if z > 10 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)
	}

}
