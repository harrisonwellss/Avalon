package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	var n int
	//var f float32
	n = 10

	var s string
	s = strconv.Itoa(n)
	fmt.Println(s)

	n, err := strconv.Atoi("abc123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	var x int64
	rand.Seed(time.Now().Unix())
	x = rand.Int63()
	s = strconv.FormatInt(x, 36)
	fmt.Println(s)

}
