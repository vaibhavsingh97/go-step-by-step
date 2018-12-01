package main

import (
	"fmt"
)

const (
	a         = 30
	b         = 30.3445
	c         = "Hello World"
	d int8    = 40
	e float32 = 40.44
	f string  = "Hey"
)

func main() {
	fmt.Println("un typed constants - compiler decides the type")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Println("Typed constants-we decide the type")
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
