package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {

	testobj := circle{
		radius: 7,
	}

	whole, decimal := math.Modf(testobj.area())

	fmt.Println(whole, decimal)

	if whole != 153 {
		t.Error("Value not within excpected Range: 153")
		t.Error("Got: ", whole)
	}

}

func TestPerimter(t *testing.T) {

	testobj := circle{
		radius: 7,
	}

	whole, decimal := math.Modf(testobj.perimeter())

	fmt.Println(whole, decimal)

	if whole != 43 {
		t.Error("Value not within excpected Range: 43")
		t.Error("Got: ", whole)
	}

}
