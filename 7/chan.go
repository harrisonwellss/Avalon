package main

import "fmt"

func combine(s []string, c chan string) {
	var str string

	for _, v := range s {
		str += v
	}
	c <- str
}

func main() {
	s := []string{"hello", "golang", "c++", "world"}

	c1 := make(chan string)
	c2 := make(chan string)

	go combine(s[:len(s)/2], c1)
	go combine(s[len(s)/2:], c2)

	//select {
	//case <-c:
	//	go combine(s[:len(s)/2], c1)
	//case <-c2:
	//	go combine(s[len(s)/2:], c2)
	//
	//}

	s1, s2 := <-c1, <-c2

	fmt.Println(s1, s2, s1+s2)

}
