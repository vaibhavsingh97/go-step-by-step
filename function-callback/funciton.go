package main

import "fmt"

func sum(slice ...int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func evenSum(f func(slice ...int) int, num ...int) int {
	var evenNumbers []int
	for _, v := range num {
		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		}

	}
	return f(evenNumbers...)
}

func oddSum(f func(slice ...int) int, num ...int) int {
	var oddNumbers []int
	for _, v := range num {
		if v%2 != 0 {
			oddNumbers = append(oddNumbers, v)
		}
	}
	return f(oddNumbers...)
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(x...))
	fmt.Println("sum of even numbers: ", evenSum(sum, x...))
	fmt.Println("sum of odd numbers: ", oddSum(sum, x...))
}
