package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "Guilherme"

	fmt.Println(strings.HasPrefix(x, "Gu"))

	fmt.Println(strings.HasSuffix(x, "rme"))
}
