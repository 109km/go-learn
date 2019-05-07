package main

import "fmt"

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
