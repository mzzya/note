package main

import "net/http"
import "log"

type Hello struct {
}

func (h Hello) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("hello world"))
}
func main() {
	err := http.ListenAndServe(":80", Hello{})
	if err != nil {
		log.Fatalf("err:%s\n", err)
	}
}
