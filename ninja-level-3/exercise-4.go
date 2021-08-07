// Hands-on exercise #4
// Create a for loop using this syntax
// for { }
// Have it print out the years you have been alive.
// solution: https://play.golang.org/p/9VpnB-I1Pz
// video: 053

package main

import "fmt"

func main() {
	i := 1995
	count := 0
	for {
		if i > 2021 {
			break
		}
		i++
		count++
	}
	fmt.Println(count)
}
