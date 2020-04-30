package main

import (
	"fmt"
)

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func dpInRecursion(coins []int, price int) int {
	if price < 0 {
		return -1
	}
	if price == 0 {
		return 0
	}

	res := price
	for i := 0; i < len(coins); i++ {
		subproblem := dpInRecursion(coins, price-coins[i])
		if subproblem == -1 {
			continue
		}
		res = min(res, 1+subproblem)
	}

	if res == price {
		res = -1
	}
	return res
}

func dpInIteration(coins []int, price int) int {
	answers := make([]int, 0)

	for index := 0; index < price+1; index++ {
		answers = append(answers, index)
	}

	for i := 0; i < len(answers); i++ {
		// 从底向上求解子问题
		for j := 0; j < len(coins); j++ {
			coin := coins[j]
			if i-coin < 0 {
				continue
			}
			// 在初始化的硬币数（最大解）及 当前价格 - 硬币面值 的解 + 1 之间进行大小比较
			// 比如价格=10的解，一定是 10-5 ， 10-2，10-1，这三个子问题解中的一个，然后加上当前这一枚硬币u
			answers[i] = min(answers[i], 1+answers[i-coin])
		}
	}

	fmt.Println(answers)
	if answers[price] == price+1 {
		return -1
	} else {
		return answers[price]
	}

}

func main() {
	coins := []int{1, 2, 5}
	price := 13

	res := dpInIteration(coins, price)
	fmt.Println(res)
}
