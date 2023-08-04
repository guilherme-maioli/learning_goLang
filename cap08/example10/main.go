package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("C:\\Users\\Guilherme\\Desktop\\Go\\cap08\\example 10\\test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		return
	}

	bs := make([]byte, stat.Size())

	_, err = file.Read(bs)

	if err != nil {
		return
	}

	str := string(bs)

	fmt.Println(str)

}
