package main

import "fmt"

func main() {
	fmt.Println("Entre com um valor a ser convertido de Fahrenheit para Celsius")

	var f float64

	fmt.Scanf("%f", &f)

	c := (f - 32) * (5. / 9.)

	fmt.Println(c)

}
