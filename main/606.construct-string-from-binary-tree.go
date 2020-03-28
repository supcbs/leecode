package main

import (
	"fmt"
	"strconv"
)

/**

题目:

你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

输入: 二叉树: [1,2,3,4]
       1
     /   \
    2     3
   /
  4

输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。
示例 2:

输入: 二叉树: [1,2,3,null,4]
       1
     /   \
    2     3
     \
      4

输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。

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

func main() {
	t1 := buildTree([]int{1, 2, 1})
	t2 := buildTree([]int{1, 1, 2})
	r := tree2str(t1, t2)
	fmt.Println(r)
}

/**
方法1:

前序遍历进行判定

时间复杂度：O(n)
空间复杂度：O(h)
*/

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	// 如果是最底层的元素则直接返回
	s := strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		return s
	}

	// 左节点不管是否为空都是需要显示的
	s += "(" + tree2str(t.Left) + ")"

	// 右节点只有不为空，才需要显示
	if t.Right != nil {
		s += "(" + tree2str(t.Right) + ")"
	}

	return s
}

