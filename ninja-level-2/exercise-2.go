// Hands-on exercise #2
// Using the following operators, write expressions and assign their values to variables:
// ==
// <=
// >=
// !=
// <
// >
// Now print each of the variables.
// solution: https://play.golang.org/p/76R-poSzaY

package main

import (
	"fmt"
)

func main() {
	a := 44
	b := 42

	fmt.Printf("%v\n", a < b)
	fmt.Printf("%v\n", a > b)
	fmt.Printf("%v\n", a == b)
	fmt.Printf("%v\n", a <= b)
	fmt.Printf("%v\n", a >= b)
	fmt.Printf("%v\n", a != b)
}
