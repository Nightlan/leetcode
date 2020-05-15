package main

import (
	"fmt"
)

func singleNumbers(nums []int) []int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	mark := ret & (-ret)
	result := []int{0, 0}
	for _, v := range nums {
		if (v & mark) == mark {
			result[0] ^= v
		} else {
			result[1] ^= v
		}
	}
	return result
}

func singleNumber(nums []int) int {
	res := 0
	for i := range nums {
		res ^= nums[i]
	}
	return res
}

func main() {
	test := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(test))
}
