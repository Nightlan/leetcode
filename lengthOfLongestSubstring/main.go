package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	res := 0
	start := 0
	tmp := 0 // 当前最大子串长度
	exist := make(map[uint8]int, len(s))
	for i := 0; i < len(s); i++ {
		pos, ok := exist[s[i]]
		if ok && pos >= start { // 出现并在子串内
			start = pos + 1 // 移动子串起始地址为之前重复地址+1
			tmp = i - pos   // 当前不重复的长度 = 当前位置 - 上一个重复字符位置
		} else {
			tmp++
		}
		if tmp > res { // 更新到最终结果
			res = tmp
		}
		exist[s[i]] = i
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
