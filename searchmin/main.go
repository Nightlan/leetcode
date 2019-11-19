package main

import "fmt"

func searchMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] >= nums[0] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main()  {
	testNums := []int{3,4,5,6,1,2}
	fmt.Println(searchMin(testNums))
}
