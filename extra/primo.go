/*
Receba um número inteiro e verifique se ele é primo ou não.
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número inteiro: ")
	fmt.Scanf("%d", &n)

	is_primo := true

	for i := 2; i < n; i++ {
		if n%i == 0 {
			is_primo = false
		}
	}

	if is_primo {
		fmt.Println("PRIMO!")
		return
	}

	fmt.Println("NÃO PRIMO!")

}
