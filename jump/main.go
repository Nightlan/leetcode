package main

import "fmt"

func jump1(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	res := 0
	for i := 0; i < len(nums); {
		nextMax := 0
		nextPos := 0
		for j := 1; j <= nums[i]; j++ {
			if i+j == len(nums)-1 {
				res++
				return res
			}
			curVal := nums[j+i] + j
			if curVal >= nextMax {
				nextMax = curVal
				nextPos = i + j
			}
		}
		i = nextPos
		res++
	}
	return res
}

func jump(nums []int) int {
	res := 0
	curMax := 0
	curEnd := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+i > curMax {
			curMax = nums[i] + i
		}
		if i == curEnd {
			curEnd = curMax
			res++
		}
	}
	return res
}

func main() {
	nums := []int{2, 1}
	fmt.Println(jump(nums))
}
