package main

import "fmt"

func main() {
	var flist []func()

	for i := 0; i < 3; i++ {
		i := i
		flist = append(flist, func() {
			//fmt.Println(&i)
			fmt.Println(i)
		})
	}

	for _, f := range flist {
		f()
	}
}
