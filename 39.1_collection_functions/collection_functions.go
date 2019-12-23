package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	res := make([]string, 0)
	for _, v := range vs {
		if f(v) == true {
			res = append(res, v)
		}
	}
	return res
}

func Map(vs []string, f func(string) string) []string {
	res := make([]string, len(vs))
	for i, v := range vs {
		res[i] = f(v)
	}
	return res
}

func main() {
	fruits := []string{"banana", "apple", "orange", "watermelon", "peach", "pear"}

	fmt.Println(Index(fruits, "apple"))
	fmt.Println(Include(fruits, "orange"))
	fmt.Println(Any(fruits, func(v string) bool {
		return strings.HasPrefix(v, "o")
	}))

	fmt.Println(All(fruits, func(v string) bool {
		return strings.HasPrefix(v, "b")
	}))

	fmt.Println(Filter(fruits, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Map(fruits, func(v string) string {
		return "Fruit:" + v
	}))

}
