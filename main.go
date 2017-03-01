package main

import (
	"fmt"
	"time"
)

func ping(in chan string, out chan string) {
	fmt.Println(<-in)
	out <- "Ping"
	time.Sleep(1 * time.Second)
}

func pong(in chan string, out chan string) {
	fmt.Println(<-in)
	out <- "Pong"
	time.Sleep(1 * time.Second)
}

func main() {
	in := make(chan string)
	out := make(chan string)
	go ping(in, out)
	go pong(out, in)

	in <- "Start"
	time.Sleep(10 * time.Second)
}
