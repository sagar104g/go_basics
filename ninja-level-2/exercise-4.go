// Hands-on exercise #4
// Write a program that
// assigns an int to a variable
// prints that int in decimal, binary, and hex
// shifts the bits of that int over 1 position to the left, and assigns that to a variable
// prints that variable in decimal, binary, and hex
// solution: https://play.golang.org/p/Ms964T8SbH
// video: 034

package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%b, %d, %#x\n", a, a, a)
	b := a << 1
	fmt.Printf("%b, %d, %#x", b, b, b)
}
