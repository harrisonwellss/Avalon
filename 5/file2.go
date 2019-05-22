package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func readd(f *os.File) (string, error) {
	var total []byte
	buf := make([]byte, 1024)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		total = append(total, buf[:n]...)
	}
	return string(total), nil
}

func main() {
	f, err := os.Open("a.txt")
	if err != nil {
		log.Fatalf("open error:%v", err)
	}

	s, err := readd(f)
	if err != nil {
		log.Fatalf("read error:%v", err)
	}
	fmt.Println(s)
}
