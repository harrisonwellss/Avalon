package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("s", " ", "separator")

//var enter = flag.String("n", "\n", "changline")
var newLine = flag.Bool("n", false, "newLine")

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))

	if *newLine {
		fmt.Println("----newline-----")
	}
	//fmt.Println(flag.Args())
}
