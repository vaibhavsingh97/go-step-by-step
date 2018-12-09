package main

import (
	"fmt"
)

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func loopFactorial(num int) int {
	total := 1
	for ; num > 0; num-- {
		total *= num
	}
	return total
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println(loopFactorial(5))
}
