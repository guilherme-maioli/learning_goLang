package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer f.Close()

	h := crc32.NewIEEE()

	_, err = io.Copy(h, f)
	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil

}

func main() {
	h1, err := getHash("teste1.txt")
	if err != nil {
		return
	}

	h2, err := getHash("teste2.txt")
	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)

}
