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

	go func() {
		time.Sleep(time.Second * 1)
		channel1 <- "channel1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		channel2 <- "channel2"
	}()

	fmt.Println("Before select:")
	// The loops' number must equal to channels' number.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
		}
	}

	fmt.Println("End of Main")
}
