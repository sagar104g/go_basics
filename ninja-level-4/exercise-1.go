// Hands-on exercise #1
// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array
// solution: https://play.golang.org/p/tD0SzV3hdf
// video: 071

package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println(v, i)
	}
}
