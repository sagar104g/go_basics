// Hands-on exercise #4
// Create and use an anonymous struct.
// solution: https://play.golang.org/p/otBHFs-lPp
// video: 089

package main

import "fmt"

func main() {
	p1 := struct {
		color     string
		door      int
		fourWheel bool
	}{
		color:     "red",
		door:      4,
		fourWheel: true,
	}

	fmt.Println(p1)
}
