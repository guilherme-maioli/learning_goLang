package main

import "fmt"

func fat(x int) int {
	if x == 0 {
		return 1
	}

	return x * fat(x-1)
}

func main() {
	num := 0
	fmt.Println("Entre com um número para calcular o fatorial: ")
	fmt.Scanf("%d", &num)

	fmt.Println("Resultado: ", fat(num))
}
