package main

import "fmt"

func maximo(x ...int) int {
	max := x[0]

	for _, value := range x[1:] {
		if value > max {
			max = value
		}
	}

	return max
}

func main() {
	fmt.Println("Escreve uma função um parametros variático e descubra o maior valor de números.")

	x := []int{1, 24, 33, 2, 4, 6, 8, 55, 88, 20}

	fmt.Println(maximo(x...))
}
