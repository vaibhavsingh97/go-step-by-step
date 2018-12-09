package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
}

type secretAgent struct {
	Person
	ltk bool
}

// any value of type secretAgent has access to method speak
// here we attach speak function to type secretAgent
// so now value of the type have access to the method speak
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		Person: Person{
			first: "adma",
			last:  "eve",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
}
