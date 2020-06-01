package main

import (
	"fmt"
	"net/http"
)

var data = []byte("哈哈哈哈哈哈")

func init() {
	for i := 0; i < 13; i++ {
		data = append(data, data...)
	}
}

func main() {
	sm := http.NewServeMux()
	sm.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		w.Write(data)
	})
	if err := http.ListenAndServe(":8185", sm); err != nil {
		fmt.Println("ListenAndServe err:", err)
	}
	http.DefaultTransport
}
