package main

import (
	"fmt"
	"time"
)

func startLoop1(src [2000][2000]int, dest [2000][2000]int) {
	i := 0
	j := 0
	for ; i < 2000; i++ {
		for ; j < 2000; j++ {
			dest[i][j] = src[i][j]
		}
	}
}
func startLoop2(src [2000][2000]int, dest [2000][2000]int) {
	i := 0
	j := 0
	for ; j < 2000; j++ {
		for ; i < 2000; i++ {
			dest[i][j] = src[i][j]
		}
	}
}

var src1 [2000][2000]int
var dest1 [2000][2000]int
var src2 [2000][2000]int
var dest2 [2000][2000]int

func main() {

	start2 := time.Now().UnixNano()
	startLoop1(src2, dest2)
	end2 := time.Now().UnixNano()
	fmt.Println((end2 - start2) / 1e6)

	start := time.Now().UnixNano()
	startLoop1(src1, dest1)
	end := time.Now().UnixNano()
	fmt.Println((end - start) / 1e6)

}
