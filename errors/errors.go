package main

import (
	"fmt"
)

func isPositiveNumber(arg int) (int, error) {
	if arg < 0 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d-%s", e.arg, e.prob)
}

func main() {
	_, e := isPositiveNumber(-1)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae)
	}
}
