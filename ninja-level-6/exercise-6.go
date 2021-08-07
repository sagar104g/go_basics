// Hands-on exercise #6
// Build and use an anonymous func
// code: https://play.golang.org/p/DQX3xEIcRe
// video: 107
package main

import "fmt"

func main() {
	func() {
		fmt.Printf("annon it is....")
	}()
}
