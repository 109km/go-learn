package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "world"

	// Close this channel, so that the `for` loop can read values from it correctly
	// If it's not closed here, the program will consider this channel is still
	// in use, so we can't read values from it
	close(c)
	for elem := range c {
		fmt.Println(elem)
	}
}
