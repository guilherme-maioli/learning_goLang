package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
)

func server(ln net.Listener) {

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
		return
	}

	c.Close()

}

func ReadMsg() {
	for {
		input := bufio.NewReader(os.Stdin)

		msg, err := input.ReadString('\n')

		if msg == "" || err != nil {
			return
		}
		client(msg)
	}
}

func main() {

	// ouve uma porta
	ln, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println(err)
		return
	}

	go server(ln)

	ReadMsg()

}
