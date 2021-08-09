// Hands-on exercise #4
// Fix the race condition you created in the previous exercise by using a mutex
// it makes sense to remove runtime.Gosched()
// code: https://github.com/GoesToEleven/go-programming
// video: 151

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
	var m sync.Mutex
	a.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			incrementer++
			lastValue = incrementer
			fmt.Println("incrementer", incrementer)
			m.Unlock()
			a.Done()
		}()
	}
	a.Wait()
	fmt.Println("lastValue", lastValue)
}
