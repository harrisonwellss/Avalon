package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Path []Point

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (path Path) Distance() float64 {
	var far float64
	for i := 0; i < len(path)-1; i++ {
		far += path[i].Distance(path[i+1])
	}
	return far
}

func (path Path) lenPoinsts() int {
	return len(path)
}

func main() {
	path := Path{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(path.Distance())
}
