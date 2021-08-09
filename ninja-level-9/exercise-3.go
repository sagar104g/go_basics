// Hands-on exercise #3
// Using goroutines, create an incrementer program
// have a variable to hold the incrementer value
// launch a bunch of goroutines
// each goroutine should
// read the incrementer value
// store it in a new variable
// yield the processor with runtime.Gosched()
// increment the new variable
// write the value in the new variable back to the incrementer variable
// use waitgroups to wait for all of your goroutines to finish
// the above will create a race condition.
// Prove that it is a race condition by using the -race flag
// if you need help, here is a hint: https://play.golang.org/p/FYGoflKQej
// code: https://github.com/GoesToEleven/go-programming
// video: 150

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("start..")
	incrementer := 0
	lastValue := 0
	var a sync.WaitGroup
	a.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			incrementer++
			lastValue = incrementer
			fmt.Println("incrementer", incrementer)
			a.Done()
		}()
	}
	a.Wait()
	fmt.Println("lastValue", lastValue)
}
