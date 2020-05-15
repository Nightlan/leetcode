package main

import "fmt"

func subarraySum1(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				res++
			}
		}
	}
	return res
}

func subarraySum(nums []int, k int) int {
	preSum := make(map[int]int, len(nums)+1)
	res := 0
	sum := 0
	preSum[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		diff := sum - k
		res += preSum[diff]
		preSum[sum]++
	}
	return res
}

func main() {
	nums := []int{1, 1, 1}
	fmt.Println(subarraySum(nums, 3))
}
