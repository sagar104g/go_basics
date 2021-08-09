// Hands-on exercise #5
// Fix the race condition you created in exercise #4 by using package atomic
// code: https://github.com/GoesToEleven/go-programming
// video: 152

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementer, 1)
			r := atomic.LoadInt64(&incrementer)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
