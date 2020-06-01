package main

import (
	"fmt"
	"sync"
	"time"
)

var a = 1

// var mx sync.Mutex

var mx sync.RWMutex

func update() {
	mx.Lock()
	a = a + 1
	fmt.Println(a)
	mx.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		go func() {
			mx.RLock()
			fmt.Println("读锁1")
			time.Sleep(time.Second * 3)
			mx.RUnlock()
		}()
	}

	for i := 0; i < 1000; i++ {
		go update()
	}
	time.Sleep(time.Second * 100)
	fmt.Println(a)
}
