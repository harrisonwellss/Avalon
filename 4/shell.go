package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	var n int
	var line string

	f := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, _ = f.ReadString('\n')
		fmt.Printf("line:#%s#", line)
		fmt.Scan(&s, &n)
		if s == "stop" {
			break
		}
		fmt.Println(s, n)
	}
}
