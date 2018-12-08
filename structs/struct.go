package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "vaibhav",
		last:  "singh",
		age:   22,
	}
	p2 := person{
		first: "admam",
		last:  "eve",
		age:   23,
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
