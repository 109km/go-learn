package main

import (
	"fmt"
	"time"
)

func main() {
	primes := sieve()

	// If we read the primes channel here,
	// this will cause channel's deadlock.
	// Because no matter reading or writing a channel,
	// this channel will be locked, so when there is no value in this channel
	// the main goroutine will wait until there is a value.
	// Unfortunately, the writing goroutine is also waiting,
	// cause the reading goroutine locked this channel.
	// So they are waiting for each other.

	// for {
	// 	fmt.Println(<-primes)
	// }

	go read(primes)
	time.Sleep(1e9)
}

func read(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 200; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()

	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
