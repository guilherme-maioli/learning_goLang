package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	h := sha1.New()
	// escrever no hasher
	h.Write([]byte("adadasd"))

	//
	bs := h.Sum([]byte{})

	fmt.Println(bs)
}
