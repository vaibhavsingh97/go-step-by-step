package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type SecretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	secretAgent := SecretAgent{
		person: person{
			first: "james",
			last:  "adman",
			age:   40,
		},
		first: "collision string",
		ltk:   true,
	}
	fmt.Println(secretAgent)
	fmt.Println(secretAgent.first, secretAgent.last, secretAgent.age, secretAgent.ltk)
	// when collision happens
	fmt.Println(secretAgent.person.first, secretAgent.person.last, secretAgent.person.age, secretAgent.first, secretAgent.ltk)
}
