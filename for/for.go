package main

import "fmt"

func main() {
	i := 0

	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 3; j < 8; j++ {
		fmt.Println(j)
	}

}
