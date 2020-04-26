package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{}
	header := result
	var bigNode *ListNode
	for {
		if l1.Val <= l2.Val {
			header.Next = l1
			l1 = l1.Next
			bigNode = l2
		} else {
			header.Next = l2
			l2 = l2.Next
			bigNode = l1
		}
		header = header.Next
		if header.Next == nil {
			header.Next = bigNode
			break
		}
	}
	return result.Next
}

func mergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	var res *ListNode
	for _, v := range lists {
		res = mergeTwoList(res, v)
	}
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}
	var tmpList []*ListNode
	size := len(lists)
	if len(lists)%2 != 0 {
		size--
		tmpList = make([]*ListNode, 0, len(lists)/2+1)
		tmpList = append(tmpList, lists[len(lists)-1])
	} else {
		tmpList = make([]*ListNode, 0, len(lists)/2)
	}
	for i := 0; i < size; i += 2 {
		tmpRes := mergeTwoList(lists[i], lists[i+1])
		tmpList = append(tmpList, tmpRes)
	}
	return mergeKLists(tmpList)
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
	n1.Next = n3
	n2.Next = n4
	lists := []*ListNode{
		n1, n2,
	}
	ret := mergeKLists(lists)
	for ret != nil {
		fmt.Println(ret.Val)
		ret = ret.Next
	}
}
