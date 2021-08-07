// Hands-on exercise #9
// A “callback” is when we pass a func into a func as an argument. For this exercise,
// pass a func into a func as an argument
// code: https://play.golang.org/p/0yGYPKh1y7
// video: 110

package main

import "fmt"

func main() {
	a := func() string {
		return fmt.Sprintf("test it is----")
	}

	foo(a)
}

func foo(x func() string) {

	fmt.Println(x())
}
