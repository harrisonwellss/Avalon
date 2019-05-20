package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var (
		x1 int8
		x2 int32
		x3 int64
		x4 uint
		x5 uint32
		x6 uint64
	)
	fmt.Println(x1, x2, x3, x4, x5, x6)

	c1 := unsafe.Sizeof(x1)
	c2 := unsafe.Sizeof(x2)
	c3 := unsafe.Sizeof(x3)
	c4 := unsafe.Sizeof(x4)
	c5 := unsafe.Sizeof(x5)
	c6 := unsafe.Sizeof(x6)

	fmt.Println(c1, c2, c3, c4, c5, c6)

	x1 = 127
	x7 := x1 + 1
	fmt.Println(x7)

}
