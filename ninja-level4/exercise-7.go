// Hands-on exercise #7
// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.
// solution: https://play.golanpackage main
// video:  077

package main

import (
	"fmt"
)

func main() {
	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	xxs := [][]string{xs1, xs2}
	for _, v := range xxs {
		fmt.Printf("%v %t\n", v, v)
	}
	fmt.Println(xxs)
	fmt.Println(len(xxs))
	fmt.Println(cap(xxs))
}
