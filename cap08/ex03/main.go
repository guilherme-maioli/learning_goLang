package main

import "fmt"

func main() {
	fmt.Println("O que é um apelido {alias} de pacote? Como criamos um?")
	fmt.Println("É uma maneira de não precisarmos sempre invocar o nome completo do pacote.")
	fmt.Println("Podemos fazer da seguinte maneira:")
	fmt.Println(`import m "math" `)
	fmt.Println("Onde `m` é o `alias` de math.")
}
