package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server(ln net.Listener) {
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}

		// trata conexão
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Recebida: ", msg)
	}

	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	go server(ln)

	var input string

	fmt.Scanln(&input)

}
