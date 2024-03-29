package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point { //конструктор
	return &Point{x, y}
}

func distance(p1, p2 *Point) float64 { // функция расчета дистации между точками
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) + math.Pow(p2.y-p1.y, 2))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	dist := distance(p1, p2)

	fmt.Printf("Расстояние между точками: %f\n", dist)
}
