package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("Mensagem de erro")

	fmt.Println(err)
}
