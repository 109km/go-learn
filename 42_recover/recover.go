package main

import "fmt"

func fail() {
	panic("An error oaccured.")
}

func test() {
	fmt.Println("test started.")
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s \n", e)
		}
	}()
	fail()
}

func main() {
	fmt.Println("Main is running.")
	test()
	fmt.Println("After panic. Main ended.")
}
