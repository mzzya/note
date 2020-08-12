package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

//openssl genrsa -out private.key 2048
//openssl rsa -in private.key -pubout -out public.key

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	privateKeyBts, err := ioutil.ReadFile("./private.key")
	if err != nil {
		panic(err)
	}
	publicKeyBts, err := ioutil.ReadFile("./public.key")
	if err != nil {
		panic(err)
	}
	privateBlock, _ := pem.Decode(privateKeyBts)
	publicBlock, _ := pem.Decode(publicKeyBts)

	privateKey, err = x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		panic(err)
	}
	publicKeyObj, err := x509.ParsePKIXPublicKey(publicBlock.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey = publicKeyObj.(*rsa.PublicKey)
}
func main() {
	eBts, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte("abcd"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("encrypt:%s\n", base64.StdEncoding.EncodeToString(eBts))

	data, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, eBts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("data:%s\n", data)
}
