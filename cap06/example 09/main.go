package main

import "fmt"

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func main() {
	defer second()

	first()

	fmt.Println("Ainda não acabou...")
}
