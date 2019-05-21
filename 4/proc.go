package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	procdir = "/proc"
	cmdline = "/cmdline"
)

func main() {
	f, err := os.Open("/proc")
	if err != nil {
		log.Fatal(err)
	}

	infos, _ := f.Readdir(-1)
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		infoname, err := strconv.Atoi(info.Name())
		if err != nil {
			continue
		}
		ioutil.ReadFile(procdir + strconv.Itoa(infoname) + cmdline)
		fmt.Println(infoname, string(filebuf))
	}
}
