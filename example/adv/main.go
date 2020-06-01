package main

import (
	"fmt"
	"math"
)

var a = 1

func main() {
	var v = 71.5150
	var v2 = 71.60
	fmt.Printf("%d\t%d\t%d\t%d\n", int(math.RoundToEven(v*100)), int(v*100), int(math.Round(71.40*100)), int(71.40*100))
	fmt.Printf("%d\t%d\t%d\t%d\n", int(math.RoundToEven(v2*100)), int(v2*100), int(math.Round(71.60*100)), int(71.60*100))
}
