package main_test

import (
	"io/ioutil"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/tidwall/gjson"
)

var data []byte

func TestMain(m *testing.M) {
	bts, err := ioutil.ReadFile("./schema.json")
	if err != nil {
		panic(err)
	}
	data = bts
	m.Run()
}

func BenchmarkGJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gjson.GetBytes(data, "type")
	}
}
func BenchmarkJsonParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonparser.GetString(data, "type")
	}
}
