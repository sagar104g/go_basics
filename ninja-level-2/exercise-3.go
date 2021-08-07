// Hands-on exercise #3
// Create TYPED and UNTYPED constants. Print the values of the constants.
// solution: https://play.golang.org/p/NutvJXWUx2
// video: 033

package main

import "fmt"

const a = 42

const b int = 43

func main() {
	// a = 44 this will give error
	fmt.Println(a, b)
}
