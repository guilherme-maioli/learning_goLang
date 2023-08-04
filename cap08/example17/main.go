package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name  string
	Idade int
}

type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Idade < b[j].Idade
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	ps := []Person{
		{"Teo", 30},
		{"Guilherme", 26},
		{"Gabrielle", 33},
		{"Thiago", 39},
	}

	sort.Sort(ByAge(ps))

	fmt.Println(ps)
}
