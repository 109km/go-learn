package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	// If the channel's value doesn't output, the worker function output won't print.
	// Because the worker function is called in goroutine's way.
	<-done // This is the signal, this line of code can stop the main routine from finishing and discard the result
}
