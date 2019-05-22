package main

import (
	"fmt"
	"strings"
)

func print() int {
	return 10
}

func smap(r rune) rune {
	fmt.Printf("%c\n", r)
	return r
}

func main() {
	s := strings.Map(func(r rune) rune {
		return r - 32
	}, "hello")
	fmt.Println(s)

}
