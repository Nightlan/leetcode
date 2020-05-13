package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	level := make([]*TreeNode, 0)
	level = append(level, root)
	for len(level) != 0 {
		nextLevel := make([]*TreeNode, 0)
		levelRes := make([]int, 0, len(level))
		for _, v := range level {
			levelRes = append(levelRes, v.Val)
			if v.Left != nil {
				nextLevel = append(nextLevel, v.Left)
			}
			if v.Right != nil {
				nextLevel = append(nextLevel, v.Right)
			}
		}
		res = append(res, levelRes)
		level = nextLevel
	}
	return res
}

func main() {

}
