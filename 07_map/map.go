package main

import "fmt"

func main() {
	var m1 = make(map[string]int)

	m1["one"] = 1
	m1["two"] = 2
	fmt.Println(m1)
	fmt.Println(len(m1))

	_, one := m1["one"] // check if this key exists

	fmt.Println(one)

	delete(m1, "one")
	fmt.Println(m1)
}
