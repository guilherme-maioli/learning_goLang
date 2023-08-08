package main

import "fmt"

func printer(ch <-chan string) {
	fmt.Println("Está função só recebe do canal ch. Isso acontece por conta da derição explicitada na assinatura da função.")
	return
}

func pinger(ch chan string) {
	fmt.Println("Está função só envia para o canal ch. Isso acontece por conta da direção explicitada na assinatura da função.")
}

func main() {
	fmt.Println("Como especificamos a direção de um tipo de canal?")

	fmt.Println("Utilizando o operador <-.")

}
