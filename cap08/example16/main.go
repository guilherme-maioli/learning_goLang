package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	kids := []Person{
		{"Ze", 9},
		{"Alberto", 10},
		{"Thomas", 9},
		{"Guilherme", 10},
	}

	sort.Sort(ByName(kids))

	fmt.Println(kids)
}
