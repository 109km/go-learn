package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

/*
	The goroutine functions will always be executed
	after the command function is executed
	But we won't know the goroutine functions executing orders.
	They ara running asynchronously.

	I guess the functions which are called in the usual way are put
	into a synchronous queue.
	And the goroutine functions will be put into an asynchronous queue.
	And the asynchronous queue won't be executed in a sequence way.

*/
func main() {

	go f("hello2")

	f("hello")

	go f("hello3")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	fmt.Scanln()
	fmt.Println("done")
}
