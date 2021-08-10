// Hands-on exercise #4
// Starting with this code, pull the values off the channel using a select statement
// solution: https://play.golang.org/p/FulKBY5JNj
// video: 167

package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c1 <-chan int, c2 chan int) {

	for {
		select {
		case v, s := <-c1:
			fmt.Println(v, s)
		case v, s := <-c2:
			fmt.Println(v, s)
			return
		}
	}
}
func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
