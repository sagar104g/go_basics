// Hands-on exercise #3
// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
// code: https://play.golang.org/p/XluEuUD0Nw
// video: 104

package main

import "fmt"

func main() {
	fmt.Println("before defer")
	defer foo1(1, 2, 3, 4, 5, 6)
	fmt.Println("after defer")
}

func foo1(x ...int) {
	sum := 0
	for v := range x {
		sum += v
	}
	fmt.Printf("%v\n", sum)
	fmt.Printf("%v, %t\n", x, x)
}
