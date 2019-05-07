package main

import "fmt"

/*
	The `Channel` is used to accept the values
	passed from goroutines which means the values can
	be passed from one thread to another.
*/

func f() {

}

func main() {

	messages := make(chan string, 2) // This channel can store two strings.

	fmt.Println("start")

	go func() {
		messages <- "ping"
		messages <- "ping2"
	}()

	go func() {
		msg_in_thread := <-messages
		fmt.Println("In the thread:", msg_in_thread)
	}()

	msg_in_main := <-messages
	fmt.Println("In the main:", msg_in_main)

	fmt.Scanln()

}
