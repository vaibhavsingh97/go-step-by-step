package main

import (
	"fmt"
)

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Vaibhav",
		last:  "singh",
		age:   21,
	}
	fmt.Println(p1)
}
