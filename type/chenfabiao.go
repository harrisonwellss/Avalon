package main

import "fmt"

func main() {
	for i := 1; i <= 9; i++ {
		fmt.Println(" ")
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
	}
}
