package main

import (
	"fmt"
	"time"
)

/*
	Select can let us wait on multiple channel operations.
*/

func main() {

	channel1 := make(chan string, 1)
	channel2 := make(chan string, 1)
	channel3 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 1)
		channel1 <- "channel1"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		channel2 <- "channel2"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		channel3 <- "channel3"
	}()

	// The loops' number must equal to channels' number.
	// If the sleep time is the same, the outputs' order is not fixed.
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
		case msg3 := <-channel3:
			fmt.Println("received", msg3)
		}
	}
	fmt.Println("main ended")
}
