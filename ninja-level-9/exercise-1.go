// Hands-on exercise #1
// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists
// code: https://github.com/GoesToEleven/go-programming
// video: 148

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("start..")
	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())
	var a sync.WaitGroup
	a.Add(2)
	go foo(&a)
	go bar(&a)
	fmt.Println("mid cpu", runtime.NumCPU())
	fmt.Println("mid gs", runtime.NumGoroutine())
	a.Wait()
	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}

func foo(a *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("foo, %v\n", i)
	}
	a.Done()
}

func bar(a *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("bar, %v\n", i)
	}
	a.Done()
}
