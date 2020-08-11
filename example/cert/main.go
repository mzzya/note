package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

var sourceBts = []byte("123456")

func main() {
	MD5()
	SHA1()
	SHA256()
	SHA512()
}

func MD5() {
	sign := md5.New()
	sign.Write(sourceBts)
	fmt.Printf("md5\t%x\n", sign.Sum(nil))
	init0 := 0x67452301
	init1 := 0xEFCDAB89
	init2 := 0x98BADCFE
	init3 := 0x10325476
	fmt.Printf("%d|%d|%d|%d\n", init0, init1, init2, init3)
}

func SHA1() {
	sign := sha1.New()
	sign.Write(sourceBts)
	fmt.Printf("sha1\t%x\n", sign.Sum(nil))
}

func SHA256() {
	sign := sha256.New()
	sign.Write(sourceBts)
	fmt.Printf("sha256\t%x\n", sign.Sum(nil))
}
func SHA512() {
	sign := sha512.New()
	sign.Write(sourceBts)
	fmt.Printf("sha512\t%x\n", sign.Sum(nil))
}

func AES() {
	// block, err := aes.NewCipher(sourceBts)
}

func RSA() {

}
