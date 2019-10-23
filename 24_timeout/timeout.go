package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "result 1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// If `default` case is set, the timeout won't work.
	// Directly goes to the end of main goroutine.
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 2")
	default:
		fmt.Println("default 2")
	}

	fmt.Println("main ended")
}
