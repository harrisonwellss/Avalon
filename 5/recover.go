package main

import (
	"fmt"
)

func printt() {
	var p *int
	fmt.Println(*p)
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	panic("不想执行下去了")
	printt()
	var i = 3

	var slice [3]int
	fmt.Println(slice[i])
}
