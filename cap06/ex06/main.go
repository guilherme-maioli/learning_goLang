package main

import "fmt"

func main() {
	fmt.Println("O que são `defer`, `panic` e `recover`? Como podemos nos recuperar de um pânico em tempo de execução? ")
	fmt.Println("São instruções que ajudam a mudar o fluxo de nossos programas.")
	fmt.Println("Defer: utilizado para que uma função ou comando seja invocado ao final da função que a pertence.")
	fmt.Println("Panic: utilizado para lançar um erro durante a execução.")
	fmt.Println("Recover: utilizado para capturar um erro durante a execução.")

	fmt.Println("Utilizamos o defer combinado com o recover para recuperar um panico. Lembrando que isso é um anti-padrão.")
}
