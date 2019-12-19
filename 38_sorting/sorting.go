package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"h", "o", "b"}
	nums := []int{23, 16, 9, 299}

	sort.Strings(strs)
	fmt.Println("Strings after sorted:", strs)

	sort.Ints(nums)
	fmt.Println("Numbers after sroted:", nums)
}
