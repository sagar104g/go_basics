// Hands-on exercise #1
// Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
// solution: https://play.golang.org/p/8BK6PXj3aG
// video: 125

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
}

func main() {
	a := person{
		First: "first1",
		Last:  "last1",
	}
	b := person{
		First: "first2",
		Last:  "last2",
	}
	c := person{
		First: "first3",
		Last:  "last3",
	}

	d := []person{a, b, c}
	fmt.Println(d)
	data, err := json.Marshal(d)
	if err != nil {
		fmt.Println("err--", err)
	}
	fmt.Println(string(data))

}
