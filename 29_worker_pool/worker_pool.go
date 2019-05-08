package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Println("worker", id, " finished job ", j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start 3 workers here
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Push values into the channel `jobs`
	// When the channel have values, the workers start to work
	for i := 1; i <= 20; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i <= 20; i++ {
		fmt.Println(<-results)
	}

}
