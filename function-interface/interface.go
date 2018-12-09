package main

import (
	"fmt"
)

type Person struct {
	first string
	last  string
}

type UnknownHuman struct {
	Person
}

func CheckType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Found: int")
	case string:
		fmt.Println("Found: string")
	case Person:
		fmt.Println("Found: Person")
	default:
		fmt.Println("Unknown Type")
	}
}

func main() {
	x := 10
	s := "Vaibhav"
	CheckType(x)
	CheckType(s)
	p1 := Person{
		first: "James",
		last:  "Bond",
	}
	CheckType(p1)
	CheckType(p1.first)
	CheckType(p1.last)
	uh1 := UnknownHuman{
		Person: Person{
			first: "vaibhav",
			last:  "Singh",
		},
	}
	CheckType(uh1)
}
