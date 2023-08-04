package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// ouve uma porta
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		// trata a conexao

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

func client() {
	// conecta ao servidor

	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Ola mundo!"

	fmt.Println("Enviando... ", msg)

	err = gob.NewEncoder(c).Encode(msg)

	if err != nil {
		fmt.Println(err)
		return
	}

	c.Close()

}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)

}
