// Hands-on exercise #7
// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
// solution: https://play.golang.org/p/IDnrJpE7vT
// video: 056

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
