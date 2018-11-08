package main

import (
	"fmt"
)

type Shape interface {
	area() float64
}

type Square int

func (sq Square) area() float64 {
	return float64(sq * sq)
}

type Rectangle struct {
	Width  int
	Height int
}

func (this Rectangle) area() float64 {
	return float64(this.Width) * float64(this.Height)
}

type Circle struct {
	Radius int
}

func (this Circle) area() float64 {
	return float64(3.14) * float64(this.Radius) * float64(this.Radius)
}

func totalArea(shapes []Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	shapes := []Shape{
		Rectangle{1, 2},
		Circle{1},
		Rectangle{1, 1},
		Square(1),
	}

	fmt.Println("Total area = ", totalArea(shapes))
}
