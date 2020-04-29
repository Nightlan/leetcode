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

func searchReversArray(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	l, r := 0, size-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if target < nums[0] && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

func main() {
	testNums := []int{3, 4, 5, 6, 1, 2}
	fmt.Println(searchMin(testNums))
	fmt.Println(searchReversArray(testNums, 2))
}
