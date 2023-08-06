package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func client(msg string) {
	// conecta ao servidor
	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enviando... ", msg)

	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()

}

func main() {

	for {
		input := bufio.NewReader(os.Stdin)

		msg, err := input.ReadString('\n')
		if err != nil || msg == "" {
			return
		}

		client(msg)
	}

}
