package Algorithm

import (
	"log"
	"math/rand"
	"time"
)

var ArrayA []int
var ArrayALength int

func init() {
	AryLen := 10
	ArrayA = make([]int, AryLen)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < AryLen; i++ {
		ArrayA[i] = rand.Intn(AryLen)
	}
	ArrayALength = len(ArrayA)
	for i, item := range ArrayA {
		if i%10 != 9 {
			log.Printf("%d\t", item)
		} else {
			log.Printf("%d\n", item)
		}
	}

}
func ShowArrayA(arrayA []int) {
	if arrayA == nil {
		arrayA = ArrayA
	}
	log.Println("*********************************************")
	for i, item := range arrayA {
		if i%10 != 9 {
			log.Printf("%d\t", item)
		} else {
			log.Printf("%d\n", item)
		}
	}
}
