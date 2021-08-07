// Hands-on exercise #2
// Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.
// solution: https://play.golang.org/p/RvvLyAMvGo
// video: 087

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
	m := map[string][]string{}
	m[p1.last] = p1.favFlavors
	m[p2.last] = p2.favFlavors

	if value, status := m[p1.last]; status {
		fmt.Println(value)
	}
	if value, status := m[p2.last]; status {
		fmt.Println(value)
	}
}
