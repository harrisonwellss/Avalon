package main

import (
	"fmt"
	"time"
)

func main() {
	var n time.Duration
	n = time.Hour
	n = 3*time.Hour + 3*time.Minute
	fmt.Println(int64(n))
	fmt.Println(n.String())
	fmt.Println(n.Seconds())
	fmt.Println(n.Minutes())

	t := time.Now()
	ti := t.Add(time.Hour)
	fmt.Println(ti)
}
