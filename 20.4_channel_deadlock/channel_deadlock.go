package main

import "fmt"

var ch chan int

func main() {
	ch = make(chan int)

	go put(ch)

	for {
		// Here will always wait for the channel's value
		// When there is no more value, this will be blocked
		// But this is in main!
		fmt.Println(<-ch)
	}
	// This line won't run
	fmt.Println("done")
}
func put(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
}
