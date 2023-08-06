package main

import (
	"example25/math"
	"fmt"
)

func main() {
	x := []float64{6.9, 9.10, 10.}

	media := math.Average(x)

	fmt.Println(media)

}
