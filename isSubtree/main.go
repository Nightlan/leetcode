package main

import (
	"container/list"
	"fmt"
)

const intMin = ^int(^uint(0) >> 1)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preTraversal(t *TreeNode) []int {
	if t == nil {
		return nil
	}
	res := make([]int, 0)
	stack := list.New()
	stack.PushBack(t)
	for stack.Len() != 0 {
		ele := stack.Back()
		stack.Remove(ele)
		node := ele.Value.(*TreeNode)
		if node == nil {
			res = append(res, intMin)
			continue
		}
		res = append(res, node.Val)
		stack.PushBack(node.Right)
		stack.PushBack(node.Left)
	}
	return res
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isSubtree1(s *TreeNode, t *TreeNode) bool {
	preS := preTraversal(s)
	preT := preTraversal(t)
	for i := 0; i < len(preS); i++ {
		end := i + len(preT)
		if end > len(preS) {
			return false
		}
		tmp := preS[i:end]
		if compare(tmp, preT) {
			return true
		}
	}
	return false
}

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

func isSubtree(s *TreeNode, t *TreeNode) bool {
	preS := preTraversal(s)
	preT := preTraversal(t)
	return kmp(preS, preT)
}

func main() {
	fmt.Println(intMin)
	t1 := &TreeNode{
		Val: 3,
	}
	t2 := &TreeNode{
		Val: 4,
	}
	t3 := &TreeNode{
		Val: 5,
	}
	t4 := &TreeNode{
		Val: 1,
	}
	t5 := &TreeNode{
		Val: 2,
	}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t6 := &TreeNode{
		Val: 4,
	}
	t7 := &TreeNode{
		Val: 1,
	}
	t8 := &TreeNode{
		Val: 2,
	}
	t6.Left = t7
	t6.Right = t8
	fmt.Println(isSubtree(t1, t6))
}
