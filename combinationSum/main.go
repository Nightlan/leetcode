package main

import (
	"fmt"
	"letcode/utils"
	"sort"
)

func sumFun(start, end, target int, candidates []int, res *utils.Stack) {
	for start <= end {
		res.Push(candidates[start])
		sum := res.Sum()
		if sum > target {
			res.Pop()
			return
		} else if sum == target{
			fmt.Println(res.String())
			res.Pop()
			return
		} else {
			sumFun(start, end, target, candidates, res)
			res.Pop()
			start++
		}
	}
}

func combinationSum(candidates []int, target int) {
	sort.Ints(candidates)
	res := utils.NewStack(0)
	start := 0
	end := len(candidates) -1
	sumFun(start, end, target, candidates, res)
}

func main(){
	testNums := []int {2,3,6,7}
	target := 7
	combinationSum(testNums, target)
	return
}
