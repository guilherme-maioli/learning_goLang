package main

import (
	"ex04/math"
	"fmt"
)

func main() {
	x := []float64{
		2.5,
		3.5,
		5.6,
		10.5,
	}

	media := math.Average(x)

	min := math.Min(x)

	max := math.Max(x)

	fmt.Println("Media: ", media)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)

}
