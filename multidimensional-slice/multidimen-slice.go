package main

import (
	"fmt"
)

func main() {
	x := []string{"adam", "chocolate", "brownie"}
	fmt.Println(x)
	y := []string{"eve", "strawberry", "hazelnut"}
	fmt.Println(y)
	z := [][]string{x, y}
	fmt.Println(z)
}
