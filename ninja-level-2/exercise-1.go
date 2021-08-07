// Hands-on exercise #1
// Write a program that prints a number in decimal, binary, and hex
// solution: https://play.golang.org/p/bAQxcEuK8O
// video: 031

package main

import (
	"fmt"
)

func main() {
	b := 42
	fmt.Printf("%d, %b, %#x", b, b, b)

}
