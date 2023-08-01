package main

import "fmt"

func par(x int) (int, bool) {
	metade := x / 2
	ehPar := x%2 == 0

	return metade, ehPar
}

func main() {
	fmt.Println("Escreva uma função que aceite um inteiro e calcule sua metade, e devolva verdadeiro caso seja par e falso se for impar:")
	numero := 0

	fmt.Println("Digite o número: ")
	fmt.Scanf("%d", &numero)

	fmt.Println(par(numero))
}
