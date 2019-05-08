package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	counter := 0
	func() {
		for t := range ticker.C {

			if counter > 9 {
				break
			}

			fmt.Println("Tick at", t)
			counter++
		}
	}()

	// time.Sleep(2 * time.Second)
	ticker.Stop()
}
