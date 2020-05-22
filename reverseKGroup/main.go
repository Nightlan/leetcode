package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
修改版：
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么将已经翻转的附加到数组最后。
示例：
给你这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->5->4
*/

func reverseKGroup1(head *ListNode, k int) *ListNode {
	res := &ListNode{}
	lastGroupTail := res
	i := 0
	var curHead *ListNode
	var curTail *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		if i%k == 0 {
			curTail = tmp
			curTail.Next = nil
		} else {
			tmp.Next = curHead
		}
		curHead = tmp
		if i%k == k-1 {
			lastGroupTail.Next = curHead
			lastGroupTail = curTail
			curHead = nil
		}
		i++
	}
	lastGroupTail.Next = curHead
	return res.Next
}

func reverse(head *ListNode) *ListNode {
	res := &ListNode{}
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = res.Next
		res.Next = tmp
	}
	return res.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	res := &ListNode{}
	lastGroupTail := res
	i := 0
	var curHead *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		if i%k == 0 {
			curHead = tmp
		}
		if i%k == k-1 {
			tmp.Next = nil
			revHead := reverse(curHead)
			lastGroupTail.Next = revHead
			lastGroupTail = curHead
			curHead = nil
		}
		i++
	}
	lastGroupTail.Next = curHead
	return res.Next
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
	n5 := &ListNode{
		Val: 5,
	}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	for node := reverseKGroup(n1, 3); node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
