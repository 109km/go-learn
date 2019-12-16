package main

import (
	"fmt"
	"time"
)

/**
 * In golang, loop1 is also much faster than loop2,
 * just because the `i,j`'s order is different.
 */

func startLoop1(src [2000][2000]int, dest [2000][2000]int) {
	start := time.Now().UnixNano()
	i := 0
	j := 0
	for ; i < 2000; i++ {
		for ; j < 2000; j++ {
			dest[i][j] = src[i][j]
		}
	}
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1e3)
}
func startLoop2(src [2000][2000]int, dest [2000][2000]int) {
	start := time.Now().UnixNano()
	i := 0
	j := 0
	for ; j < 2000; j++ {
		for ; i < 2000; i++ {
			dest[i][j] = src[i][j]
		}
	}
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1e3)
}

var src1 [2000][2000]int
var dest1 [2000][2000]int
var src2 [2000][2000]int
var dest2 [2000][2000]int

func main() {
	startLoop1(src1, dest1)
	startLoop2(src2, dest2)
}
