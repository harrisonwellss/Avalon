package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	fmt.Println(buf[:n])

	r := bufio.NewReader(f)
	r.ReadByte()
}
