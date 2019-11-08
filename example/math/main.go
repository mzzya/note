package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print(math.Log2(1000))
	//fmt.Println(math.E, math.Exp(1), math.Exp2(1), math.Expm1(1), math.Log10(906)/math.Log10(8087+906), math.Atan(906)*2/math.Pi)
	// fmt.Println(1 / (1 + math.Exp(3)))
	// fmt.Println(1 / (1 + math.Expm1(3)))
	// fmt.Println(math.Log(2000), math.Atan(300))
	log(10)
	//atan(1000)
}

func log(max int) {
	for i := 0; i <= max; i++ {
		ll := math.Log10(float64(i)) / math.Log10(float64(1000))
		fmt.Printf("%d\t%f\t%f\n", i, math.Log10(float64(i)), ll)
	}
}
func atan(max int) {
	for i := 0; i <= max; i++ {
		fmt.Printf("%d\t%f\n", i, math.Atan(float64(i))*2/math.Pi)
	}
}
