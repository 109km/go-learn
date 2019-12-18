package main

import (
	"fmt"
	"time"
)

var ch chan int

func main() {
	ch = make(chan int)
	go put(ch)

	fmt.Println("From main:", <-ch) // Only prints 0

	go out(ch) // Outputs 1 - 99.
	time.Sleep(time.Millisecond)

}

func out(channel chan int) {
	for {
		fmt.Println(<-channel)
	}
}

func put(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
}
