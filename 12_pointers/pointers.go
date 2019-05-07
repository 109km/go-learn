package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeropointer(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1

	zeroval(i)
	fmt.Println(i)

	iptr := &i
	zeropointer(iptr)

	fmt.Println(i)
	fmt.Println(&i)
}
