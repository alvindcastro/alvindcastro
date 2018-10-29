package main

import (
	"fmt"
)

var y = 43
var z int

func main() {
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
