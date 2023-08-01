package main

import "fmt"

func media(x []float64) float64 {
	total := 0.0

	for _, v := range x {
		total += v
	}

	return total / float64(len(x))
}

func main() {
	x := []float64{98, 93, 77, 82, 83}

	fmt.Println("Média: ", media(x))
}
