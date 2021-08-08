// Hands-on exercise #1
// Create a value and assign it to a variable.
// Print the address of that value.
// code: https://play.golang.org/p/57kW8xd0qT
// video: 116

package main

import "fmt"

func main() {
	a := 42
	b := &a

	fmt.Println(a, b)

}
