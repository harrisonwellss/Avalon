package main

import "fmt"

func fib(num int) int {
	if num <= 2 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

func main() {
	for i := 0; i < 100; i++ {
		nums := fib(i)

		fmt.Println(i)
		fmt.Println(nums)
	}
}
