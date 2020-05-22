package main

import "fmt"

func findTheLongestSubstring(s string) int {
	lastStatus := make([]int, 1<<5)
	res := 0
	status := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if lastStatus[status] != 0 || status == 0 {
			curVal := i + 1 - lastStatus[status]
			if curVal > res {
				res = curVal
			}
		} else {
			lastStatus[status] = i + 1
		}
	}
	return res
}

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}
