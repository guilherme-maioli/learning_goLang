package main

import (
	"fmt"
	"time"
)

func Sleep(t int) {
	select {
	case <-time.After(time.Second * time.Duration(t)):
		return
	}
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		Sleep(2)
	}

	var input string
	fmt.Scanln(&input)
}
