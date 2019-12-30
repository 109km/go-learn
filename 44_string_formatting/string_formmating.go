package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var log = fmt.Printf

func main() {
	p := &point{1, 2}
	log("%v\n", p)
	log("%+v\n", p)
	log("%#v\n", p)

	log("%T\n", p)
	log("%t\n", false)
	log("%d\n", 123)
	log("%b\n", 10)
	log("%x\n", 10)
	log("%f\n", 3.14159265357)

	log("%p\n", &p)
	log("|%8d|%8s|\n", 1898892, "Hello")

	test := fmt.Sprintf("%8s|%8s", "This is ", "a string\n")
	log(test)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
