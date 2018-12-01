package main

import (
	"fmt"
	"runtime"
)

var x1 int8 = -128

func main() {
	x := 10
	y := 10.45786
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Println(x1)
	fmt.Printf("%T\n", x1)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.Compiler)

}
