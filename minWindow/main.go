package main

import "fmt"

func check(t, c map[uint8]int) bool {
	for k, v := range t {
		if c[k] < v {
			return false
		}
	}
	return true
}

func minWindow(s string, t string) string {
	targetChar := make(map[uint8]int, len(t))
	curChar := make(map[uint8]int, len(t))
	for i := 0; i < len(t); i++ {
		targetChar[t[i]]++
		curChar[t[i]] = 0
	}
	res := ""
	l := 0
	r := 0
	for i := 0; i < len(s); i++ {
		_, ok := curChar[s[i]]
		if ok {
			curChar[s[i]]++
			r = i + 1
		}
		if check(targetChar, curChar) {
			for j := l; l < r; j++ {
				if r-j < len(res) || len(res) == 0 {
					res = s[j:r]
				}
				_, ok = curChar[s[j]]
				if ok {
					curChar[s[j]]--
					l = j + 1
					if curChar[s[j]] < targetChar[s[j]] {
						break
					}
				}
			}
		}
	}
	return res
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
