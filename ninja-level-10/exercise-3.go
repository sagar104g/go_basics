// Hands-on exercise #3
// Starting with this code, pull the values off the channel using a for range loop
// solution: https://play.golang.org/p/D3N4Tq54SN
// video: 166

package main

import (
	"fmt"
)

func main() {
	c := gen()
	// var wg sync.WaitGroup
	// wg.Add(1)
	receive(c)
	// go receive(c, &wg)
	// wg.Wait()
	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) { //, wg *sync.WaitGroup

	for i := range c {
		fmt.Println(i)
	}
	// wg.Done()
}
