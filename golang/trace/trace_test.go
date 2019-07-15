package main

import "testing"

type Student struct {
	ID   uint
	Name string
}

var count = 100000000

func BenchmarkForRange(b *testing.B) {
	var students = make([]Student, count)
	for k, item := range students {
		if item.ID%2 == 0 {
			students[k].Name = "aaaaaaa"
		} else {
			students[k].Name = "bbbbbbb"
		}
	}
}

func BenchmarkFor(b *testing.B) {
	var students = make([]Student, count)
	for i := 0; i < len(students); i++ {
		if students[i].ID%2 == 0 {
			students[i].Name = "aaaaaaa"
		} else {
			students[i].Name = "bbbbbbb"
		}
	}
}
