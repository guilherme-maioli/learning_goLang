package main

import (
	"fmt"
	"math/rand"
)

func gera_pacote(N, n int) []int {
	pacote := []int{}

	for i := 1; i <= n; i++ {
		pacote = append(pacote, rand.Intn(500)+1)
	}

	return pacote
}

func simula_album(N, n int) int {
	album := map[int]int{}
	pacotes := 0
	for len(album) < N {
		pacote := gera_pacote(N, n)
		for _, v := range pacote {
			album[v]++
		}
		pacotes++
	}

	return pacotes
}

func media(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total / len(x)
}

func main() {
	N := 500
	n := 4

	resultados := []int{}

	for i := 1; i <= 1000; i++ {
		resultados = append(resultados, simula_album(N, n))
	}

	fmt.Println("Média de pacotes comprados: ", media(resultados))
}
