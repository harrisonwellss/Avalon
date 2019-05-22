package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("aaa")
	}
	fmt.Println(string(out))
}
