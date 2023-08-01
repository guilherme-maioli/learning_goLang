package main

import "fmt"

func main() {

	// tamanho da fatia 3, capacidade 9
	slice := make([]int, 3, 9)

	fmt.Println("Tamanho da fatia é: ", len(slice))
}
