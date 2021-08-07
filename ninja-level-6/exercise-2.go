// Hands-on exercise #2
// create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in
// code: https://play.golang.org/p/B0yRxtBQPD
// video: 103

package main

import "fmt"

func main() {
	foo1(1, 2, 3, 4, 5, 6)
	a := []int{1, 2, 3, 4, 5, 6}
	foo1(a...)
	bar1(a)
}

func foo1(x ...int) {
	sum := 0
	for v := range x {
		sum += v
	}
	fmt.Printf("%v\n", sum)
	fmt.Printf("%v, %t\n", x, x)
}

func bar1(x []int) {
	sum := 0
	for v := range x {
		sum += v
	}
	fmt.Printf("%v\n", sum)
	fmt.Printf("%v, %t\n", x, x)
}
