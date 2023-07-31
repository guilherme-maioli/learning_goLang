package main

import "fmt"

func main() {
	fmt.Println("Escreve um programa que exiba números de 1 a 100, mas para múltiplos de 3, mostre `Fizz` no lugar do número e para multiplos de 5 exiba `buzz`. Para números que seja múltiplos tanto de 3 quanto de 5 mostre `FizzBuzz`")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
