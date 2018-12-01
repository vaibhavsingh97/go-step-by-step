package main

import (
	"fmt"
)

func main() {
	s := "Hello World"
	fmt.Println(s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	fmt.Printf("\n")
	for i, j := range s {
		fmt.Println(i, j)
	}
}
