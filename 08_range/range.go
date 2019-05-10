package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 5}
	sum := 0

	/*
		When using `for ... range` grammar,
		no matter what type of the range iterator is,
		the first variable will always be the `key`,
		and the second variable is the `value`.
		So if we don't need the `key` we can use `_` instead of the
		variable's name
	*/
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	maps := map[string]int{"one": 1, "two": 2}
	mapsSum := 0
	for k := range maps {
		fmt.Println(k)
	}
	for _, v := range maps {
		mapsSum += v
	}
	fmt.Println(mapsSum)
}
