package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string) {
	for i := 0; i < 3; i++ {
		cakeName := "S cake " + strconv.Itoa(i)
		fmt.Println("Making a cake and sending ... ", cakeName)
		cs <- cakeName
	}
}

func receiveCakeAndPack(cs chan string) {
	for i := 0; i < 3; i++ {
		s := <-cs
		fmt.Println("packing and received ... ", s)
	}
}
func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs)
	go receiveCakeAndPack(cs)

	time.Sleep(4 * time.Second)
}
