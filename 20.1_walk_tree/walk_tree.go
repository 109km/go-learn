package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

func recWalk(t *tree.Tree, ch chan int) {
	if t != nil {
		recWalk(t.Left, ch)
		ch <- t.Value
		recWalk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		x1, ok1 := <-ch1
		x2, ok2 := <-ch2
		switch {
		case ok1 != ok2:
			return false
		case !ok1:
			return true
		case x1 != x2:
			return false
		default:
		}
	}

}

func main() {
	ch := make(chan int)
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		ret := Same(tree.New(1), tree.New(1))
		ch1 <- ret
	}()
	go func() {
		ret := Same(tree.New(1), tree.New(2))
		ch2 <- ret
	}()
	go Walk(tree.New(1), ch)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

	for v := range ch {
		fmt.Println(v)
	}

}
