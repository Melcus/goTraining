package main

import (
	"fmt"
)

// Package level
var x int = 42

func main() {
	fmt.Println(x)
	foo()

	//  block level
	y := 10
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
