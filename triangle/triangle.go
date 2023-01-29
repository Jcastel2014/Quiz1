package main

import "math"

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

}
