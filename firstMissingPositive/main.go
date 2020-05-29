package main

import "fmt"

/*
41. 缺失的第一个正数
给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。



示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1


提示：

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的额外空间。
*/

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	isExistOne := false
	for i := range nums {
		if nums[i] == 1 {
			isExistOne = true
			break
		}
	}
	if !isExistOne {
		return 1
	}
	if len(nums) == 1 {
		return 2
	}
	mark := len(nums) + 1
	for i := range nums {
		if nums[i] >= mark || nums[i] <= 0 {
			nums[i] = 1
		}
	}
	for i := range nums {
		val := nums[i]
		if val < 0 {
			val = -val
		}
		val--
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}
	res := mark
	for i := range nums {
		if nums[i] > 0 {
			res = i + 1
			break
		}
	}
	return res
}

func main() {
	nums := []int{0, -1, 3, 1}
	fmt.Println(firstMissingPositive(nums))
}
