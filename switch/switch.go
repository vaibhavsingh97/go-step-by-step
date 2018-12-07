package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("This statement will not print")
	case (2 == 4):
		fmt.Println("This is not true, so will not print")
	case (3 == 3):
		fmt.Println("This will print :)")
		// This keyword helps doesn't break after case is matched,
		// also goes to below cases
		fallthrough
	case (4 == 4):
		fmt.Println("Also, true, but does it print?")
		fallthrough
	case (5 == 4):
		fmt.Println("this is not true")
		fallthrough
	case (7 == 7):
		fmt.Println("True, 7 == 7")
		// because of fallthrough, default statement is also execute
		fallthrough
	default:
		fmt.Println("This is a default case")
	}

	fmt.Println("=========Next example==========")
	x := "hello"
	switch x {
	case "test":
		fmt.Println("This is a test")
	case "hello":
		fmt.Println("hello World")
	case "test1":
		fmt.Println("This is another test")
	default:
		fmt.Println("None of these cases matched")
	}
}
