package main

import (
	"fmt"
	"unsafe"
)

func main() {
	q2 := [...]int{4: 2, 10: -1}

	fmt.Println(len(q2))

	a1 := [3]int{1, 2, 3}

	var a2 [3]int

	a2 = a1

	fmt.Println(&a1[0], &a2[0])
	fmt.Println(unsafe.Sizeof(a1))

	a2[1] = 99

	fmt.Println(a1, a2)
}
