package main

import "fmt"

type rect struct {
	width  int
	height int
}

// Add methods to struct `rect`
// Pass the value by address.
func (r *rect) area() int {
	r.width = 15 // Can mutate the `r` outside.
	return r.width * r.height
}

// Pass the value by value
func (r rect) permi() int {
	r.width = 5 // Can't mutate the `r` outside.
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	rptr := &r
	fmt.Println(r.area())
	fmt.Println(r.permi())

	fmt.Println(rptr.area())
	fmt.Println(rptr.permi())

	fmt.Println(r, rptr)
}
