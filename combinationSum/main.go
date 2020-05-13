package main

import (
	"fmt"
	"letcode/utils"
	"sort"
)

func sumFun(start, end, target int, candidates []int, stack *utils.Stack, res *[][]int) {
	for start <= end {
		stack.Push(candidates[start])
		sum := stack.Sum()
		if sum > target {
			stack.Pop()
			return
		} else if sum == target {
			sRes := stack.Buf()
			tmp := make([]int, len(sRes))
			copy(tmp, sRes)
			*res = append(*res, tmp)
			stack.Pop()
			return
		} else {
			sumFun(start, end, target, candidates, stack, res)
			stack.Pop()
			start++
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	stack := utils.NewStack(0)
	start := 0
	end := len(candidates) - 1
	res := make([][]int, 0)
	sumFun(start, end, target, candidates, stack, &res)
	return res
}

func main() {
	testNums := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(testNums, target))
	return
}
