package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan bool)

	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
		channel <- true
	}()

	<-channel

	// stop2 := timer2.Stop()
	// if stop2 {
	// 	fmt.Println("Timer2 stopped")
	// }
}
