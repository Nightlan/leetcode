package main

import "fmt"

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
通过

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func doPermute(nums []int, curRes []int, res *[][]int) {
	for i, v := range nums {
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

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	doPermute(nums, nil, &res)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums)
	for _, v := range res {
		fmt.Println(v)
	}
}
