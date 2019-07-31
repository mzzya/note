package main

import (
	"log"
	"time"
)

type student struct {
	ID   int
	Name string
}

func main() {
	t, _ := time.Now().MarshalJSON()
	log.Printf("%s\n", t)
	var stu = student{}
	St(&stu)
	println(stu == student{})
	s := make([]struct{}, 10)
	print(s)
	a := 1
	b := 2
	c := Add(a, b)
	print(c)
}

func Add(a int, b int) int {
	return a + b
}

func St(s *student) {
	return
}
