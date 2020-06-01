package main

import (
	"fmt"
	"runtime"
)

var a = 10

func update() {
	a = 1
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
}

func main() {

	for i := 0; i < 10; i++ {
		go update()
	}
	update()
	runtime.Goexit()
	fmt.Println(runtime.GOMAXPROCS(4))
	fmt.Println("NumCPU", runtime.NumCPU())
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
}
