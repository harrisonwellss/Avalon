package main

import (
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

	buf := make([]byte, 10124)
	n, _ := f.Read(buf)
	fmt.Printf("%d\n###%s###\n", n, string(buf[:n]))

}
