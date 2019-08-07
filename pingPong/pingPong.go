package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go Ping(ch)
	go Pong(ch)
	ch <- "Start"
	time.Sleep(15 * time.Second)
}

func Ping(ch chan string) {
	for {
		fmt.Println("# :", <-ch)
		time.Sleep(1 * time.Second)
		ch <- "Ping"
	}
}

func Pong(ch chan string) {
	for {
		fmt.Println("@ :", <-ch)
		time.Sleep(1 * time.Second)
		ch <- "Pong"
	}
}
