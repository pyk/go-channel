package main

import (
	"fmt"
)

func NewServer(wait, done chan bool) {
	fmt.Println("SERVER EXECUTED")
	wait <- true

	for {
		select {
		case <-done:
			fmt.Println("SERVER DONE")
			return
		default:
			fmt.Println("NewServer loop")
		}
	}
}

func main() {
	fmt.Println("MAIN EXECUTED")
	wait := make(chan bool)
	done := make(chan bool)
	go NewServer(wait, done)
	<-wait
	fmt.Println("MAIN already receive signal that server are listening")
	done <- true
	fmt.Println("MAIN EXIT")
}
