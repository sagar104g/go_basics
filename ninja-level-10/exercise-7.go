// Hands-on exercise #7
// write a program that
// launches 10 goroutines
// each goroutine adds 10 numbers to a channel
// pull the numbers off the channel and print them
// solutions:
// https://play.golang.org/p/R-zqsKS03P
// https://play.golang.org/p/quWnlwzs2z
// student solutions:
// https://twitter.com/mannion
// https://gist.github.com/mannion007/3c8899913974c1027ef6f13ec37b2b3f
// video: 170

package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

func gen(x, y int) <-chan int {
	c := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				c <- j
			}
		}()
		fmt.Println("ROUTINES", runtime.NumGoroutine())
	}
	return c
}
