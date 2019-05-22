package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type Student struct {
	Id   int
	Name string
}

func add(args []string) error {
	//name := args[0]
	//id := args[1]
	fmt.Println("call add")
	fmt.Println("args", args)
	return nil
}

func list(args []string) error {
	return errors.New("list error")
}

func main() {
	artionmap := map[string]func([]string) error{
		"add":  add,
		"list": list,
	}

	f := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _ := f.ReadString('\n')
		line = strings.TrimSpace(line)
		args := strings.Fields(line)
		fmt.Println(args)
		if len(args) == 0 {
			continue
		}

		cmd := args[0]
		args = args[1:]
		actionfunc := artionmap[cmd]

		if actionfunc == nil {
			fmt.Println("bad cmd", cmd)
			continue
		}
		fmt.Println("found cmd func")
		continue

		err := actionfunc(args)
		if err != nil {
			fmt.Printf("execute action %s error:%s\n", cmd, err)
			continue
		}
	}
}
