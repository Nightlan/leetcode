package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{
		Next: head,
	}
	lastNode := res
	var noRepeatedNode *ListNode
	repeated := false
	for head != nil {
		if head.Next != nil {
			if head.Val == head.Next.Val {
				if !repeated {
					repeated = true
					noRepeatedNode = lastNode
				}
			} else {
				if repeated {
					repeated = false
					noRepeatedNode.Next = head.Next
					lastNode = noRepeatedNode
				} else {
					lastNode = head
				}
			}
		} else {
			if repeated {
				noRepeatedNode.Next = nil
			}
		}
		head = head.Next
	}
	return res.Next
}

func main() {
	n1 := &ListNode{
		Val: 1,
	}
	n2 := &ListNode{
		Val: 2,
	}
	n1.Next = n2
	n3 := &ListNode{
		Val: 3,
	}
	n2.Next = n3
	n4 := &ListNode{
		Val: 3,
	}
	n3.Next = n4
	n5 := &ListNode{
		Val: 4,
	}
	n4.Next = n5
	n6 := &ListNode{
		Val: 4,
	}
	n5.Next = n6
	n7 := &ListNode{
		Val: 5,
	}
	n6.Next = n7
	for head := deleteDuplicates(n1); head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
