// Hands-on exercise #1
// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
// solution:
// https://play.golang.org/p/ouRHmH_POg
// video: 086

package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p1 := person{
		first:      "sagar",
		last:       "yadav",
		favFlavors: []string{"a", "b"},
	}
	p2 := person{
		first:      "sagar1",
		last:       "yadav1",
		favFlavors: []string{"a1", "b1"},
	}
	fmt.Println(p1, p2)
}
