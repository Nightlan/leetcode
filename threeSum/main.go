package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func doThreeSum(nums []int, ele []int, sum int, res *[][]int) {
	if len(ele) == 3 || sum > 0 {
		return
	}
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := make([]int, 0, len(ele)+1)
		tmp = append(tmp, ele...)
		tmp = append(tmp, nums[i])
		tmpSum := sum + nums[i]
		if tmpSum == 0 && len(tmp) == 3 {
			*res = append(*res, tmp)
			return
		}
		doThreeSum(nums[i+1:], tmp, tmpSum, res)
	}
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	doThreeSum(nums, nil, 0, &res)
	return res
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for first := range nums {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -nums[first]
		third := len(nums) - 1
		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third > second && nums[second]+nums[third] > target {
				third--
			}
			if third == second {
				break
			}
			if nums[second]+nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
