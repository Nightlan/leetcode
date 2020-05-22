package main

import (
	"fmt"
	"time"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree1(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	rootVal := preOrder[0]
	index := make(map[int]int, 0)
	for i, v := range inOrder {
		index[v] = i
	}
	root := &TreeNode{
		Val: rootVal,
	}
	pos := index[rootVal]
	root.Left = buildTree1(preOrder[1:pos+1], inOrder[:pos])
	root.Right = buildTree1(preOrder[1+pos:], inOrder[pos+1:])
	return root
}

func doBuildTree(preOrder []int, pl, pr int, inOrder []int, il, ir int, index map[int]int) *TreeNode {
	if pl > pr {
		return nil
	}
	rootVal := preOrder[pl]
	root := &TreeNode{
		Val: rootVal,
	}
	pos := index[rootVal]
	lTreeSize := pos - il
	root.Left = doBuildTree(preOrder, pl+1, pl+lTreeSize, inOrder, il, pos-1, index)
	root.Right = doBuildTree(preOrder, pl+lTreeSize+1, pr, inOrder, pos+1, ir, index)
	return root
}

func buildTree(preOrder []int, inOrder []int) *TreeNode {
	index := make(map[int]int, 0)
	for i, v := range inOrder {
		index[v] = i
	}
	return doBuildTree(preOrder, 0, len(preOrder)-1, inOrder, 0, len(inOrder)-1, index)
}

func main() {
	preOrder := []int{3, 9, 20, 15, 7}
	inOrder := []int{9, 3, 15, 20, 7}
	nowTime := time.Now()
	buildTree(preOrder, inOrder)
	fmt.Println(time.Now().Sub(nowTime).Nanoseconds())
}
