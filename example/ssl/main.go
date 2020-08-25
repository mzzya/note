package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	cli := http.Client{}
	resp, err := cli.Get("https://testbapi.colipu.com/v")
	if err != nil {
		panic(err)
	}
	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("resp:%s\n", bts)
	time.Sleep(time.Minute)
}
