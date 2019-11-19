package main

import (
	"fmt"
	"letcode/utils"
	"sort"
)

func numCount(candidates []int) map[int]int {
	countMap := make(map[int]int, len(candidates))
	for i := range candidates {
		countMap[candidates[i]]++
	}
	return countMap
}

func sumFun(start, end, target int, candidates []int, res *utils.Stack, countMap map[int]int) {
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
			sumFun(start+1, end, target, candidates, res, countMap)
			res.Pop()
			start++
		}
	}
}

func combinationSum(candidates []int, target int) {
	sort.Ints(candidates)
	res := utils.NewStack(len(candidates))
	start := 0
	end := len(candidates) -1
	countMap := numCount(candidates)
	sumFun(start, end, target, candidates, res, countMap)
}

func main(){
	testNums := []int {10, 1,2,7,6,1,5}
	target := 8
	combinationSum(testNums, target)
	return
}
