package main

import "fmt"

func main() {
	fmt.Println(`Sabemos que na base 10 o maior número de um digito é 9 e o maior numero de dois dígito é 99 
	(...) qual o maior número de oito dígitos? `)

	fmt.Println(`Resposta: base 10 = 10^8 - 1 = `, 100000000-1)
	fmt.Println(`Resposta: base 2 = 2^8 - 1 = `, 2*2*2*2*2*2*2*2-1)
}
