// Hands-on exercise #3
// Create a for loop using this syntax
// for condition { }
// Have it print out the years you have been alive.
// solution: https://play.golang.org/p/tnyqBPJ-i5
// video: 052

package main

import "fmt"

func main() {
	i := 1995
	count := 0
	for i < 2021 {
		i++
		count++
	}
	fmt.Println(count)
}
