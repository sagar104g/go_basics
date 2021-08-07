// Hands-on exercise #4
// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person
// code: https://play.golang.org/p/NnXyWdqbbw
// video: 105

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	a := person{
		first: "first",
		last:  "last",
		age:   4,
	}
	a.foo1(1, 2, 3, 4, 5)
}

func (p person) foo1(x ...int) {

	fmt.Printf("%v, %t\n", x, x)
	fmt.Printf("%v\n", p)
}
