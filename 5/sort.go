package main

import (
	"fmt"
	"sort"
)

type Students struct {
	Id   int
	Name string
}

func main() {
	s := []int{2, 3, 1, 5, 9, 7}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	ss := []Students{}
	ss = append(ss, Students{
		Name: "aa",
		Id:   2,
	})

	ss = append(ss, Students{
		Name: "bb",
		Id:   3,
	})

	ss = append(ss, Students{
		Name: "cc",
		Id:   1,
	})

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Name < ss[j].Name
	})

	fmt.Println(ss)
}
