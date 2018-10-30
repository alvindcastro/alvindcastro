package main

import (
	"fmt"
)

type burger int

var x burger

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)

}
