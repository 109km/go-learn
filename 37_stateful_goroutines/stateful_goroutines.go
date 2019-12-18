package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {
	var readCount uint64
	var writeCount uint64
	readChannel := make(chan readOp)
	writeChannel := make(chan writeOp)

	// Block 1
	// Start a goroutine to wait for the channel
	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-readChannel:
				// Put the state's value into the read.resp channel
				read.resp <- state[read.key]
			case write := <-writeChannel:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	// Start 100 reading operations' goroutines
	for r := 0; r < 100; r++ {
		go func() {
			for {
				// Create a read operation
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				// Put this read operation into the read channel
				readChannel <- read
				// Once the read.resp has a value,
				// means the reading operation is completed in the block 1,
				// then we can add one to `readCount`
				<-read.resp
				atomic.AddUint64(&readCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Start 10 writing operations
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				writeChannel <- write
				// Once write.resp has a true value,
				// means the writing operation in block 1 is finished.
				<-write.resp
				atomic.AddUint64(&writeCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readCount)
	fmt.Println("readOps:", readOpsFinal)

	writeOpsFinal := atomic.LoadUint64(&writeCount)
	fmt.Println("writeOps:", writeOpsFinal)
}
