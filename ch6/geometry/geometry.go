package main

import (
	"math"
	"fmt"
)

type Point struct {
	x, y float64
}

type Path []Point

func main() {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x - q.x, p.y - q.y)
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i - 1].Distance(path[i])
		}
	}
	return sum
}



