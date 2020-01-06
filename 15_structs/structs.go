package main

import "fmt"

type vertex struct {
	flag    bool
	counter int16
	pi      float32
}

var v1 vertex

func main() {

	v2 := &vertex{}
	fmt.Printf("%+v\n", v1)
	fmt.Printf("%+v\n", *v2)
}
