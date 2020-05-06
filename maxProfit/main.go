package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	buy := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > buy {
			tmpRes := prices[i] - buy
			if tmpRes > res {
				res = tmpRes
			}
		}
		if prices[i] < buy {
			buy = prices[i]
		}
	}
	return res
}

func main() {
	prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit(prices))
}
