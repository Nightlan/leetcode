package main

import "fmt"

func kmp(a, b []int) bool {
	next := make([]int, len(b))
	next[0] = -1
	for i, j := 0, -1; i < len(b)-1; {
		if j == -1 || b[j] == b[i] {
			j++
			i++
			next[i] = j
		} else {
			j = next[j]
		}
	}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if j == -1 || a[i] == b[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}
	if j == len(b) {
		return true
	}
	return false
}

func main() {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1}
	b := []int{1, 1, 1, 1, 2, 2}
	fmt.Println(kmp(a, b))
}
