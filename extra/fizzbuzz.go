/*
Print números de 1 a 100.
Quando o número for divisivel por 3, print FIZZ ao invés do número
Quando o número for divisivel por 5, print BUZZ ao invés do número
Quando o número for divisivel por 3 e 5, print FIZZBUZZ ao invés do número
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var msg string
	for i := 1; i <= 100; i++ {
		msg = ""

		if i%3 == 0 {
			msg = msg + "FIZZ"
		}
		if i%5 == 0 {
			msg = msg + "BUZZ"
		}
		if msg == "" {
			msg = strconv.Itoa(i)
		}

		fmt.Println(msg)
	}

}
