package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"adam":    12121212121212,
		"eve":     22323232323232323,
		"vaibhav": 9898989889898989,
	}
	fmt.Println(m)
	fmt.Println(m["adam"])
	fmt.Println(m["test"]) // if there is no value is stored in a key then it return a 0
	v, ok := m["test"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["eve"]; ok {
		fmt.Println("This is if value", v)
	}

	//Adding nee element to map
	m["john"] = 909090909090
	fmt.Println(m)

	for key, value := range m {
		fmt.Println(key, value)
	}

	// Delete key and value
	delete(m, "vaibhav")
	fmt.Println(m)

	// If key doesn't exist then it will not give error while deleting
	delete(m, "random key")
	fmt.Println(m)
}
