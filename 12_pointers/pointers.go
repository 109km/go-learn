package main

import "fmt"

type user struct {
	name string
	age  int
}

func zeroval(ival int) {
	ival = 0
}

func zeropointer(iptr *int) {
	*iptr = 0
}
func escaspeToHeap() *user {
	u := user{
		name: "Hello",
		age:  10,
	}
	return &u
}
func main() {

	i := 1
	j := 2
	l := 3

	zeroval(i)
	fmt.Println(i) // Still 1

	iptr := &i
	zeropointer(iptr)

	user1 := escaspeToHeap()

	fmt.Printf("%p\n", user1)

	fmt.Println(&i)
	fmt.Println(&j)
	fmt.Println(&l)
}
