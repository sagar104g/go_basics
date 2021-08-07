// Hands-on exercise #10
// Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:
// code: https://play.golang.org/p/a56uWtoFSL
// video: 111

package main

import "fmt"

func main() {

	a := foo()
	b := foo()

	a(1)
	a(1)
	a(1)
	b(2)
	b(2)
	b(2)

}

func foo() func(x int) {
	b := 0
	return func(x int) {
		b += x
		fmt.Println(b)
	}
}
