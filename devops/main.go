package main

import "sync/atomic"

var a uint

func bbb()

func main() {
	atomic.StoreUintptr(&a, 1)
	bbb()()
}
