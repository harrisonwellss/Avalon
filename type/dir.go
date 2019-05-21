package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//dirName := flag.String("d", ".", "dir name")
	//flag.Parse()
	//
	//f, err := os.Open(*dirName)

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	infos, _ := f.Readdir(-1)
	fmt.Println(infos)
	for _, info := range infos {
		fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
	}
}
