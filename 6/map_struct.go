package main

import "fmt"

type Student struct {
	Id   int
	Name string
}

func main() {
	m := make(map[string]*Student)
	m["binggan"] = &Student{
		Id:   1,
		Name: "binggan",
	}
	m["binggan"].Id = 2

	fmt.Println(m)

	m1 := make(map[string]*int)
	n := 1
	m1["a"] = &n

	fmt.Println(m1)
}
