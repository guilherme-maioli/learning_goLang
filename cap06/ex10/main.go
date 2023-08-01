package main

import "fmt"

func square(x *float64) {
	*x = *x * *x
}

func main() {
	fmt.Println("Qual é o valor de x após a execução desse programa?")
	x := 1.5

	square(&x)

	fmt.Println(x)

}
