package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	result := &ListNode{}
	curRet := result
	for head != nil {
		next := head.Next
		if next != nil {
			tmpHead := next.Next
			curRet.Next = next
			next.Next = head
			curRet = head
			head = tmpHead
			curRet.Next = nil
		} else {
			curRet.Next = head
			head = head.Next
		}
	}
	return result.Next
}

func main() {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n3 := &ListNode{
		Val: 3,
	}
	n4 := &ListNode{
		Val: 4,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	ret := swapPairs(n1)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
