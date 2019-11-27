package main

import (
	"fmt"
)

func maxSubArray(array []int) (res int) {
	if len(array) == 0 {
		return
	}
	res = array[0]
	sum := 0
	for _, v := range array {
		if sum <= 0 {
			sum = v
		} else {
			sum += v
		}
		if sum > res {
			res = sum
		}
	}
	return
}

func main() {
	testArray := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(testArray))
	return
}
