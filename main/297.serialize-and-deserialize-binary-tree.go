package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**

题目:

序列化是将一个数据结构或者对象转换为连续的比特位的操作，
进而可以将转换后的数据存储在一个文件或者内存中，
同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。
这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

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
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = nums[0]

	ch := make(chan *TreeNode, len(nums))
	ch <- root
	nums = nums[1:]

	for i := 0; i < len(nums); i++ {
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val: nums[i],
			}

			ch <- tree.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val: nums[i],
			}

			ch <- tree.Right
		}
	}

	return root
}

func printTree(t *TreeNode) {
	ret := make([]int,0)
	stack := make([]*TreeNode,0)
	stack = append(stack, t)
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		ret = append(ret, cur.Val)

		if cur.Left != nil {
			stack = append(stack,cur.Left)
		}

		if cur.Right != nil {
			stack = append(stack,cur.Right)
		}
	}

	fmt.Println(ret)
}

func main() {
	c := Constructor()

	str := "1,2,X,X,3,4,X,X,5,X,X,"
	root := c.deserialize(str)
	printTree(root)

	//t1 := buildTree([]int{1, 2, 3,-1,-1,4,5})
	//string := c.serialize(t1)
	//fmt.Println(string)
}

/**
方法1:

整体专为数组字符串的方式来进行操作
不管是组成字符串还是组成树，都是前序遍历

时间复杂度：O(n)
空间复杂度：O(h)
*/

type nodes []string

func (n *nodes) buildString(root *TreeNode) {
	// 如果为空节点，则说明该分支递归完了，加上X代表结束
	if root == nil {
		*n = append(*n, "X", ",")
		return
	}

	// 添加字符串和逗号
	*n = append(*n, strconv.Itoa(root.Val), ",")
	n.buildString(root.Left)
	n.buildString(root.Right)
}

func (n *nodes) buildTree() *TreeNode {
	// 判断字符串长度和是否是结束标识
	if len(*n) > 0 && (*n)[0] == "X" {
		return nil
	}

	// 将第一位转为数字
	data, _ := strconv.Atoi((*n)[0])

	// 构建树
	t := &TreeNode{data, nil, nil}

	// 剔除掉一个元素之后继续构建树的左节点
	*n = (*n)[1:]
	t.Left = n.buildTree()

	// 继续剔除一个元素之后继续构建树的右节点
	*n = (*n)[1:]
	t.Right = n.buildTree()

	return t
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	buffer := nodes([]string{})
	(&buffer).buildString(root)
	return strings.Join(buffer, "")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	dataSlice := strings.Split(data, ",")
	n := nodes(dataSlice)
	return (&n).buildTree()
}
