package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, inorderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal(root.Right)...)
	return ret
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)
	return ret
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	ret = append(ret, postorderTraversal(root.Left)...)
	ret = append(ret, postorderTraversal(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}

func preorderTraversal1(root *TreeNode) []int {
	var res []int
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			res = append(res, root.Val)
			root = root.Left
		}
		backEle := stack.Back()
		stack.Remove(backEle)
		back := backEle.Value.(*TreeNode)
		root = back.Right
	}
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		backEle := stack.Back()
		stack.Remove(backEle)
		root = backEle.Value.(*TreeNode)
		res = append(res, root.Val)
		if root.Right != nil {
			stack.PushBack(root.Right)
		}
		if root.Left != nil {
			stack.PushBack(root.Left)
		}
	}
	return res
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := list.New()
	tmpRes := list.New()
	stack.PushBack(root)
	for stack.Len() != 0 {
		backEle := stack.Back()
		stack.Remove(backEle)
		root = backEle.Value.(*TreeNode)
		tmpRes.PushFront(root.Val)
		if root.Left != nil {
			stack.PushBack(root.Left)
		}
		if root.Right != nil {
			stack.PushBack(root.Right)
		}
	}
	res := make([]int, 0, tmpRes.Len())
	for v := tmpRes.Front(); v != nil; v = v.Next() {
		res = append(res, v.Value.(int))
	}
	return res
}

func inorderTraversal1(root *TreeNode) []int {
	var ret []int
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}
		backEle := stack.Back()
		stack.Remove(backEle)
		root = backEle.Value.(*TreeNode)
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

func main() {

}
