// Hands-on exercise #5
// Print out the remainder (modulus) which is found for each number between 10 and 100 when it is divided by 4.
// solution: https://play.golang.org/p/ohfJOW9euy
// video: 054

package main

import "fmt"

func main() {
	for i := 10; i < 100; i++ {
		fmt.Printf("num %v = modulo %v\n", i, i%4)
	}
}
