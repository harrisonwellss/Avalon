package main

import "fmt"

func main() {
	var a, b int
	a = 10
	b = 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	a = a + 3 // a+=3
	if a >= b {
		fmt.Println(true)
	}

	s := "hellp"
	s += "world"

	if (a > b && b > 3) || b > 10 {
		fmt.Println("aslo true")
	}

	if a == b || a != b {

	}

}
