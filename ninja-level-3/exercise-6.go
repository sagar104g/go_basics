// Hands-on exercise #6
// Create a program that shows the “if statement” in action.
// solution: https://play.golang.org/p/DpZ_FLfn5s
// video: 055

package main

import (
	"fmt"
)

func main() {
	i := 3
	j := 2
	if j > i {
		fmt.Println("i > j")
	} else if j == i {
		fmt.Println("i == j")
	} else if j < i {
		fmt.Println("j < i")
	}
}
