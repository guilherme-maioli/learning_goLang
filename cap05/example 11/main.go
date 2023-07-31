package main

import "fmt"

func main() {
	elementos := map[string]string{

		"H": "Hidrogênio",
		"O": "Oxigênio",
		"C": "Carbono",
		"N": "Nitrogênio",
	}

	fmt.Println(elementos["O"])

	name, ok := elementos["Gui"]

	fmt.Println(name, ok)

	if name, ok := elementos["Teste"]; ok {
		fmt.Println(name)
	}
}
