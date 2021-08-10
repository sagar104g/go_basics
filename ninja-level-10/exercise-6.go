// Hands-on exercise #6
// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them
// solution: https://play.golang.org/p/096Lk1BR7o
// video: 169

package main

import "fmt"

func main() {
	fmt.Println("starting...")
	c1 := make(chan int)
	c2 := make(chan int)
	input(c1, c2)
	output(c1, c2)
	fmt.Println("finish...")
}

func input(c1, c2 chan<- int) {
	go func() {
		for i := 0; i < 100; i++ {
			c1 <- i
		}
		c2 <- 1
		close(c1)
		close(c2)
	}()
}

func output(c1, c2 <-chan int) {
	for {
		select {
		case i, s := <-c1:
			fmt.Println(i, s)
		case i, s := <-c2:
			fmt.Println(i, s)
			return
		}
	}
}
