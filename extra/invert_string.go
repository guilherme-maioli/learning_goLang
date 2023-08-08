/*
Faça um programa com uma função que recebe uma frase. Para cada palavra nesta frase, inverta a ordem das letras. Exiba o resultado.
*/

package main

import (
	"fmt"
	"strings"
)

func invert_string(msg string) string {
	var msg_invert string

	a := strings.Split(msg, " ")
	for _, v := range a {
		a_invert := ""
		for _, r := range v {
			a_invert = string(r) + a_invert
		}
		msg_invert = msg_invert + a_invert + " "
	}

	return msg_invert
}

func main() {
	msg := "Está é a frase original"

	fmt.Println(invert_string(msg))

}
