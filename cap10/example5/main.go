package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {

	for i := 1; ; i++ {
		c <- fmt.Sprintf("%d Pinger", i)
	}
}
func ponger(c chan string) {

	for i := 1; ; i++ {
		c <- fmt.Sprintf("%d Ponger", i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}

}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string

	fmt.Scanln(&input)
}
