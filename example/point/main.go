package main

import (
	"fmt"
	"time"
)

var global = 1

func main() {
	fmt.Printf("global:%p\n", &global)
	time.Sleep(time.Hour)
}
