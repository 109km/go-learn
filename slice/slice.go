package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:5]

	s = append(s, 20)
	s = append(s, 21)

	fmt.Println(s)

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	fmt.Println("cpy:", c[2:5])
	fmt.Println("cpy:", c[:2])
	fmt.Println("cpy:", c[5:])

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoD[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("twoD:", twoD)

}
