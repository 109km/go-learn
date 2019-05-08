package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(500 * time.Millisecond)
	for req := range requests {
		<-limiter
		fmt.Println("request tick ", req, time.Now())
	}

	bussLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		bussLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			bussLimiter <- t
		}
	}()

	bussRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		bussRequests <- i
	}
	close(bussRequests)
	for req := range bussRequests {
		<-bussLimiter
		fmt.Println("request buss ", req, time.Now())
	}
}
