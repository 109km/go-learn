package main

import "fmt"

/*
The `channel` type have two ways of use:
1. Store the value in to the channel
2. Read the value from the channel

The `channel` is more like a queue: first in first out.

In here, the directions means we can declare how we are
going to use the channel.

*/

func ping(pings chan<- string, msg string) {
	fmt.Println("msg is stored into pings")
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	fmt.Println("msg is read from pings")
	msg := <-pings
	fmt.Println("msg is stored into pongs")
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hello,world")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
