package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

/*
序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例:

你可以将以下二叉树：

    1
   / \
  2   3
     / \
    4   5

序列化为 "[1,2,3,null,null,4,5]"
提示: 这与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

说明: 不要使用类的成员 / 全局 / 静态变量来存储状态，你的序列化和反序列化算法应该是无状态的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	seq := list.New()
	seq.PushBack(root)
	var buffer strings.Builder
	for seq.Len() != 0 {
		ele := seq.Front()
		seq.Remove(ele)
		node := ele.Value.(*TreeNode)
		if node != nil {
			buffer.WriteString(strconv.Itoa(node.Val) + ",")
			seq.PushBack(node.Left)
			seq.PushBack(node.Right)
		} else {
			buffer.WriteString("null,")
		}
	}
	res := buffer.String()
	return res[:len(res)-1]
}

func (this *Codec) getNode(str string) *TreeNode {
	if str == "null" {
		return nil
	}
	val, err := strconv.Atoi(str)
	if err != nil {
		return nil
	}
	return &TreeNode{
		Val: val,
	}
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	res := this.getNode(nodes[0])
	seq := list.New()
	isLeft := true
	parent := res
	for i := 1; i < len(nodes); i++ {
		node := this.getNode(nodes[i])
		if isLeft {
			parent.Left = node
		} else {
			parent.Right = node
		}
		if node != nil {
			seq.PushBack(node)
		}
		isLeft = !isLeft
		if isLeft {
			ele := seq.Front()
			if ele != nil {
				seq.Remove(ele)
				parent = ele.Value.(*TreeNode)
			}
		}
	}
	return res
}

func main() {
	t1 := &TreeNode{
		Val: 1,
	}
	t2 := &TreeNode{
		Val: 2,
	}
	t3 := &TreeNode{
		Val: 3,
	}
	t4 := &TreeNode{
		Val: 4,
	}
	t5 := &TreeNode{
		Val: 5,
	}
	t1.Left = t2
	t1.Right = t3
	t3.Left = t4
	t3.Right = t5
	code := Codec{}
	str := code.serialize(t1)
	fmt.Println(str)
	root := code.deserialize(str)
	fmt.Println(code.serialize(root))
}
