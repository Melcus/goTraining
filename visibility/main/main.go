package main

import (
	"fmt"

	"github.com/Melcus/visibility/factory"
)

func main() {

	fmt.Println(factory.UnoriginalVarName) // Bubel
	// fmt.Println(factory.derpString)   lower case, NOT VALID
	factory.PrintPrivateVal() // Using an accessor
}
