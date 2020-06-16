package main

import (
	"container/list"
	"fmt"
)

/*
根据每日 气温 列表，请重新生成一个列表，对应位置的输出是需要再等待多久温度才会升高超过该日的天数。如果之后都不会升高，请在该位置用 0 来代替。

例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。

提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/daily-temperatures
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	stack := list.New()
	stack.PushBack(0)
	for i := 1; i < len(T); i++ {
		for stack.Len() != 0 {
			ele := stack.Back()
			index := ele.Value.(int)
			if T[i] > T[index] {
				res[index] = i - index
				stack.Remove(ele)
			} else {
				break
			}
		}
		stack.PushBack(i)
	}
	return res
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
