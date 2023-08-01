package main

import "fmt"

func main() {
	elementos := map[string]map[string]string{
		"H": {
			"nome":   "Hidrogênio",
			"estado": "gasoso",
		},
		"He": {
			"nome":   "Hélio",
			"estado": "gasoso",
		},
		"Li": {
			"nome":   "Lítio",
			"estado": "sólido",
		},
		"Be": {
			"nome":   "Beryllium",
			"estado": "Sólido",
		},
		"B": {
			"nome":   "Bório",
			"estado": "sólido",
		},
		"C": {
			"nome":   "Carbono",
			"estado": "Sólido",
		},
		"N": {
			"nome":   "Nitrogênio",
			"estado": "gasoso",
		},
		"O": {
			"nome":   "Oxigênio",
			"estado": "gasoso",
		},
		"F": {
			"nome":   "Fluor",
			"estado": "gasoso",
		},
		"Ne": {
			"nome":   "Neon",
			"estado": "gasoso",
		},
	}

	var elemento string

	fmt.Println("Entre com a sigla de um elemento: ")
	fmt.Scanf("%s", &elemento)

	if el, ok := elementos[elemento]; ok {
		fmt.Println(el["nome"], ":", el["estado"])
	}
}
