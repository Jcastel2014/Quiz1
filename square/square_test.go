package main

import "testing"

func TestSquare(t *testing.T) {
	a, p := square(7)

	if a != 49 || p != 28 {
		t.Error("Unexpected Results")
	}
}
