package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%d: ping", i)
	}
}

func printer(c chan string) {
	for {
		// 1 forma
		//msg := <-c
		//fmt.Println(msg)

		// 2 forma
		fmt.Println(<-c)

		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}
