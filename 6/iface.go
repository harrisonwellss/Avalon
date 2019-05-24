package main

import "fmt"

type IInstance interface {
	Instance() float64
}

type Path []Point

func (p Path) Instance() float64 {

}

func main() {
	var i IInstance
	p := Path{{1, 2}, {3, 4}}
	i = p
	fmt.Println(i)
}
