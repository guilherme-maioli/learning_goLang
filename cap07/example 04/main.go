package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Olá meu nome é", p.Name)
}

type Androide struct {
	Person
	Model string
}

func main() {
	meuAndroide := Androide{}
	meuAndroide.Person.Talk()
}
