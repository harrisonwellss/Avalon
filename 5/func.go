package main

import (
	"fmt"
	"os"
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi(a, b int) int {
	return a * b
}

func main() {
	funcmap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"x": multi,
	}

	m, _ := strconv.Atoi(os.Args[1])
	n, _ := strconv.Atoi(os.Args[3])

	f := funcmap[os.Args[2]]
	if f != nil {
		fmt.Println(f(m, n))
	}
}
