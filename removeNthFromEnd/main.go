package main

import "fmt"

/*

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	forward := head
	nthEnd := head
	var nthLast *ListNode
	res := head
	for forward != nil {
		forward = forward.Next
		if n > 0 {
			n--
		} else {
			nthLast = nthEnd
			nthEnd = nthEnd.Next
		}
	}
	if nthLast != nil {
		nthLast.Next = nthEnd.Next
	} else {
		res = res.Next
	}
	return res
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
	for node := removeNthFromEnd(n1, 4); node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
