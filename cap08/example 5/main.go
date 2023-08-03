package main

import (
	"fmt"
	"strings"
)

func main() {

	nomes := []string{"Guilherme", "Adriana", "Carlos", "Gabi"}

	fmt.Println(strings.Join(nomes, "+"))

}
