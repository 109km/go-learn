package main

import "fmt"

/*
	The `Channel` is used to accept the values
	passed from goroutines which means the values can
	be passed from one thread to another.
*/

func main() {

	msg := make(chan string)         // This type of grammar declares block channel
	messages := make(chan string, 2) // This non-block channel, it can store two strings.
	msg2 := make(chan string)
	fmt.Println("start")

	go func() {
		msg <- "msg"
		msg2 <- "msg2"
		messages <- "ping"
		messages <- "ping2"
	}()

	// If the channel is a blocking channel,
	// it must be received in a single coroutine
	go func() {
		<-msg
	}()

	go func() {
		<-msg2
	}()

	go func() {
		msg_in_thread := <-messages
		fmt.Println("In the thread:", msg_in_thread)
	}()

	msg_in_main := <-messages
	fmt.Println("In the main:", msg_in_main)

	fmt.Scanln()

}
