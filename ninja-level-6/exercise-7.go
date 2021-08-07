// Hands-on exercise #7
// Assign a func to a variable, then call that func
// code: https://play.golang.org/p/_Qu7ZAyFDH
// video: 108
package main

import "fmt"

func main() {
	a := func() {
		fmt.Printf("variable it is....")
	}

	a()
}
