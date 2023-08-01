package main

import "fmt"

func main() {
	fmt.Println("Como podemos acessar o quarto elemento de um array de uma fatia?")

	x := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Quarto elemento: ", x[3])

}
