package main

import (
	"fmt"
	"os"
)

func main() {

	m := os.Args[1]
	n := os.Args[3]

	switch os.Args[2] {
	case "+":
		fmt.Println(m + n)
		//case "-":
		//	fmt.Println(m - n)

	}

}
