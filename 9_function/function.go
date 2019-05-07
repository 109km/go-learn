package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func multiple(a int, b int) int {
	return a * b
}

// Multiple return values
func plusAndMinus(a int, b int) (int, int) {
	return a + b, a - b
}

// Variadic function
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("plus", plus(2, 3))
	fmt.Println("multiple", multiple(3, 5))
	fmt.Println(plusAndMinus(5, 8))
	fmt.Println("sum", sum(1, 2, 3, 4, 5))

	totalNumbers := []int{10, 11, 12}
	fmt.Println("sum array:", sum(totalNumbers...))
}
