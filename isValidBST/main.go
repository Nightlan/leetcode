package main

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	stack := list.New()
	lastVal := 0
	isFirst := true
	for stack.Len() != 0 || root != nil {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		ele := stack.Back()
		stack.Remove(ele)
		node := ele.Value.(*TreeNode)
		if !isFirst {
			if node.Val <= lastVal {
				return false
			}
		}
		isFirst = false
		lastVal = node.Val
		root = node.Right
	}
	return true
}

func main() {

}
