package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {
	addr, err := net.LookupAddr("testclpcrmapi.colipu.com")
	if err != nil {
		fmt.Printf("LookupAddr error:%s\n", err)
	}
	fmt.Printf("addraddr,%#v", addr)
	return
	cli := http.Client{}
	resp, err := cli.Get("http://testclpcrmapi.colipu.com/api/v1/Opportunities/Search?currentUserId=112787&customerCode=&PageIndex=1")
	if err != nil {
		fmt.Printf("do error:%s\n", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read error:%s\n", err)
	}
	fmt.Printf("result:%s\n", body)
}
