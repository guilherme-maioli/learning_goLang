package main

import (
	"os"
)

func main() {

	file, err := os.Create("C:\\Users\\Guilherme\\Desktop\\Go\\cap08\\example 11\\test.txt")

	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString("Testando escrita em arquivo!")

}
