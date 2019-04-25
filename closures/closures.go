package main

import "fmt"

func autoPlus(gap int) func() int {
	var i int = 0
	return func() int {
		i = i + gap
		return i
	}
}

func main() {
	var tick = autoPlus(1)
	fmt.Println("add 1 per time", tick())
	fmt.Println("add 1 per time", tick())

	var anotherTick = autoPlus(2)
	fmt.Println("add 2 per time:", anotherTick())
	fmt.Println("add 2 per time:", anotherTick())
}
