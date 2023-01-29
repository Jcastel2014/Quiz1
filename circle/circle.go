package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	roundy := circle{
		radius: 9,
	}
	fmt.Println("Radius of the circle is ", roundy.radius)

	fmt.Println("Area of a circle is ", roundy.area())
	fmt.Println("Circumference of a circle is ", roundy.perimeter())

}
