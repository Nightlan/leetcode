package main

import "fmt"

func doDecodeString(s string, pos *int) []byte {
	count := 0
	res := make([]byte, 0)
	for *pos < len(s) {
		if s[*pos] >= '0' && s[*pos] <= '9' {
			count = (int(s[*pos]) - '0') + count*10
		} else if s[*pos] == '[' {
			*pos++
			tmp := doDecodeString(s, pos)
			for i := 0; i < count; i++ {
				res = append(res, tmp...)
			}
			count = 0
		} else if s[*pos] == ']' {
			break
		} else {
			res = append(res, s[*pos])
		}
		*pos++
	}
	return res
}

func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}
	start := 0
	return string(doDecodeString(s, &start))
}

func main() {
	fmt.Println(decodeString("100[leetcode]"))
}
