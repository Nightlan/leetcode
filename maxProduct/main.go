package main

import "fmt"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	curMax := nums[0]
	curMin := nums[0]
	for i := 1; i < len(nums); i++ {
		tmpMax := nums[i] * curMax
		tmpMin := nums[i] * curMin
		if tmpMax > tmpMin {
			curMax = tmpMax
			curMin = tmpMin
		} else {
			curMax = tmpMin
			curMin = tmpMax
		}
		if nums[i] < curMin {
			curMin = nums[i]
		}
		if nums[i] > curMax {
			curMax = nums[i]
		}
		if curMax > res {
			res = curMax
		}
	}
	return res
}

func main() {
	nums := []int{-2, 3, -4}
	fmt.Println(maxProduct(nums))
}
