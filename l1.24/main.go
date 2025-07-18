package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dx, dy := p.x-other.x, p.y-other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	exmpP1, exmpP2 := NewPoint(3.4, 8.7), NewPoint(1.2, 4.5)

	distance := exmpP1.Distance(exmpP2)
	fmt.Printf("Результат расстояния: %f\n", distance)
}
