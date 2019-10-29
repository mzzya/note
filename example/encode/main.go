package main

import (
	"encoding/hex"
	"fmt"
	"unicode/utf8"
)

func main() {

	a, size := utf8.DecodeRuneInString("我来个去")
	by, err := hex.DecodeString("我")
	fmt.Printf("%b\t%#X\t len:%d\n", a, a, size)
	fmt.Printf("%b\t%#X\n%s", by, by, err)
}
