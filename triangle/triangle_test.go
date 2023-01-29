package main

import (
	"math"
	"testing"
)

func TestTrianglearea(t *testing.T) {

	rightAngleTriangle := triangle{
		base:   50,
		height: 60,
	}

	if rightAngleTriangle.area() != (50*60)/2 {
		t.Error("UnEqual Results", rightAngleTriangle.area(), " does not Equal", (50*60)/2)

	}

}

func TestTriangleperimeter(t *testing.T) {
	rightAngleTriangle := triangle{
		base:   50,
		height: 60,
	}

	perimiter := (math.Round(rightAngleTriangle.perimiter()*100) / 100)

	if perimiter != 188.1 {
		t.Error("Unexpected Return: ", perimiter)
	}

}
