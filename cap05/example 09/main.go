package main

import "fmt"

func main() {
	// todas as chaves do dicionário é STRING
	// todos os valores relacionados a chave é INTEIRO
	x := make(map[string]int)
	x["key"] = 10

	fmt.Println(x)
}
