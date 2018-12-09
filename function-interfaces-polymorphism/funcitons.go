package main

import (
	"fmt"
)

// A value can be of more than two type

type Person struct {
	first string
	last  string
}

type SecretAgent struct {
	Person
	ltk bool
}

// If "person" and "secret agent" has SPEAK method than they are of type human
type human interface {
	speak()
}

func bar(h human) {
	// doing assetion - confirming the type
	switch h.(type) {
	case Person:
		fmt.Println("I was passed to the bar function-SWITCH", h.(Person).first)
	case SecretAgent:
		fmt.Println("I was passed to the bar function-SWITCH", h.(SecretAgent).first)
	}
	fmt.Println("I was passed to the bar function", h)
}

func (s SecretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "-- the secret agent speak")
}

func (p Person) speak() {
	fmt.Println("I am", p.first, p.last, " -- the person speak")
}

func main() {
	sa1 := SecretAgent{
		Person: Person{
			first: "Adma",
			last:  "Eve",
		},
		ltk: true,
	}
	p1 := Person{
		first: "James",
		last:  "bond",
	}
	fmt.Println(sa1)
	sa1.speak()

	fmt.Println(p1)
	bar(sa1)
	bar(p1)
}
