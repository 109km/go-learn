package main

import "fmt"

func main() {
	a := 2
	b := 5
	if b%a == 0 {
		fmt.Println("This is an even number")
	} else {
		fmt.Println("This is an odd number")
	}

	if num := 10; num < 9 {
		fmt.Println(num, "is smaller than 9")
	} else if num > 9 {
		fmt.Println(num, "is larger than 9")
	} else {
		fmt.Println(num, "is eaqual to 9")
	}
}
