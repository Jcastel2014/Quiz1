package main

import (
	"fmt"
	"math"
)

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return (t.base * t.height) / 2

}

func (t triangle) perimiter() float64 {
	hyp := math.Sqrt((t.base * t.base) * (t.height * t.height))

	return hyp + t.base + t.height

}
func main() {

	rightAngleTriangle := triangle{
		base:   20,
		height: 40,
	}

	fmt.Println("Triangle has a base of: ", rightAngleTriangle.base)
	fmt.Println("Triangle has a height of: ", rightAngleTriangle.height)

	fmt.Println("Area of the triangle is", rightAngleTriangle.area())
	fmt.Println("Perimeter of the triangle is", rightAngleTriangle.perimiter())

}
