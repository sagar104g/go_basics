// Hands-on exercise #9
// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
// solution: https://play.golang.org/p/h-8FnjpzWB
// video: 058

package main

import "fmt"

func main() {
	name := "favSport"
	switch name {
	case "favSport1":
		fmt.Println("favSport1")
	case "favSport2":
		fmt.Println("favSport2")
	case "favSport":
		fmt.Println("favSport")
	default:
		fmt.Println("default")
	}
}
