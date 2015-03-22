package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("pinger: send ping to c")
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		fmt.Println("ponger: send ping to c")
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		fmt.Println("printer: waiting c")
		msg := <-c
		fmt.Println("printer: c received")
		fmt.Println(msg)
		fmt.Println("DONE")
		time.Sleep(time.Second * 1)
	}
}
func main() {
	c := make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
