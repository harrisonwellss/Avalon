package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("starting print file~~~ %d\n", os.Args[i])
		printFIle(os.Args[i])
	}
}

func printFIle(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}
