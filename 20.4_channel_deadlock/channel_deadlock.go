package main

import "fmt"

var ch chan int

func main() {
	ch = make(chan int)

	go put(ch)
	// Here will always wait for the channel's value
	// When there is no more value, this will be blocked
	// But this is in main!
	for i := range ch {
		fmt.Println(i)
	}
	// This line won't run
	fmt.Println("done")
}
func put(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	// If we close the channel manually,
	// the `for-range` will detect the channel is closed or not,
	// so it won't cause deadlock.
	// close(ch)
}
