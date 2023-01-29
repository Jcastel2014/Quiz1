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

	per := math.Hypot(t.base, t.height) + t.base + t.height

	return per

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
