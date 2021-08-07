// Hands-on exercise #8
// Create a func which returns a func
// assign the returned func to a variable
// call the returned func
// code: https://play.golang.org/p/c2HwqVE1Rd
// video: 109

package main

import "fmt"

func main() {
	a := foo()

	fmt.Println(a())
}

func foo() func() string {
	return func() string {
		return fmt.Sprintf("test it is")
	}
}
