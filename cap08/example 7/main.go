package main

import (
	"fmt"
	"strings"
)

func main() {
	nome := "Guilherme"

	fmt.Println(strings.Replace(nome, "e", "a", 1))
}
