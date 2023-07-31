package main

import "fmt"

func main() {
	elementos := make(map[string]string)

	elementos["H"] = "Hidrogênio"
	elementos["O"] = "Oxigênio"
	elementos["C"] = "Carbono"
	elementos["N"] = "Nitrogênio"

	fmt.Println(elementos["O"])

	name, ok := elementos["Gui"]

	fmt.Println(name, ok)

	if name, ok := elementos["Teste"]; ok {
		fmt.Println(name)
	}
}
