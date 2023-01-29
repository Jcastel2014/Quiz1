package main

import "fmt"

func square(side float64) (float64, float64) {
	area := side * side
	perimeter := 4 * side

	return area, perimeter

}

func main() {
	area, perimeter := square(7)

	fmt.Println("Area of the square is ", area)
	fmt.Println("Perimeter of the square is ", perimeter)
}
