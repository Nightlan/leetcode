package main

import "fmt"

func palindrome(l, r int, s string) int {
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			break
		}
		l--
		r++
	}
	return r - l - 1
}

func longestPalindrome(s string) string {
	l := 0
	r := 0
	for i := 0; i < len(s); i++ {
		len1 := palindrome(i, i+1, s)
		len2 := palindrome(i-1, i+1, s)
		if len1 > len2 {
			len2 = len1
		}
		if len2 > r-l {
			r = i + len2/2 + 1
			l = i - (len2-1)/2
		}
	}
	return s[l:r]
}

func main() {
	fmt.Println(longestPalindrome("babd"))
}
