package main

import "fmt"
import "time"

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
	<-done
}
