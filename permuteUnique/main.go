package main

import (
	"fmt"
	"sort"
)

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func doPermute(nums []int, curRes []int, res *[][]int) {
	lastVal := int((^uint(0)) >> 1)
	for i, v := range nums {
		if v == lastVal {
			continue
		}
		lastVal = v
		tmpRes := make([]int, 0, len(nums))
		tmpRes = append(tmpRes, curRes...)
		tmpRes = append(tmpRes, v)
		if len(nums) == 1 {
			*res = append(*res, tmpRes)
		} else {
			condArray := make([]int, 0, len(nums)-1)
			condArray = append(condArray, nums[:i]...)
			condArray = append(condArray, nums[i+1:]...)
			doPermute(condArray, tmpRes, res)
		}
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	doPermute(nums, nil, &res)
	return res
}

func main() {
	nums := []int{-1, 2, -1, 2, 1, -1, 2, 1}
	res := permuteUnique(nums)
	for _, v := range res {
		fmt.Println(v)
	}
}
