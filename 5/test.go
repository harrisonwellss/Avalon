package main

import "fmt"

func main() {

	a := func() int { return 10 }
	fmt.Println(a())
}
