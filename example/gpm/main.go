package main

import (
	"fmt"
	"runtime"
	"sync"

	_ "go.uber.org/automaxprocs"
)

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumCgoCall())
	fmt.Println(runtime.NumGoroutine())

	// runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("A: ", i)
			wg.Done()
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()
}
