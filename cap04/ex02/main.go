package main

import "fmt"

func main() {
	fmt.Println("Escreve todos os programas de 1 a 100 que seja divisíveis por 3.")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
