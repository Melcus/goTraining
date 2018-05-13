package main

import (
	"fmt"
)

func zero(z *int) {
	*z = 0
}

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	zero(&a)
	fmt.Println(a)
}
