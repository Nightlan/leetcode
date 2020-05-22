package main

import "fmt"

func validStr(head, tail *int, s string) bool {
	for *head < *tail {
		if s[*head] != s[*tail] {
			return false
		}
		*head++
		*tail--
	}
	return true
}

func validPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	head := 0
	tail := len(s) - 1
	res := validStr(&head, &tail, s)
	if !res {
		newHead1 := head + 1
		newTail1 := tail
		newHead2 := head
		newTail2 := tail - 1
		if !validStr(&newHead1, &newTail1, s) &&
			!validStr(&newHead2, &newTail2, s) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
