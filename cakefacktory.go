package main

import (
	"fmt"
	"strconv"
	"time"
)

var i int

func makeCakeAndSend(cs chan string) {
	i = i + 1
	cakeName := "Strawberry Cake " + strconv.Itoa(i)
	fmt.Println("Making cake and sending ...", cakeName)
	cs <- cakeName
}

func receiveCakeAndPack(cs chan string) {
	s := <-cs
	fmt.Println("Packing received cake: ", s)
}

func main() {
	cs := make(chan string)
	for i := 0; i < 3; i++ {
		go makeCakeAndSend(cs)
		go receiveCakeAndPack(cs)

		time.Sleep(4 * time.Second)
	}
}
