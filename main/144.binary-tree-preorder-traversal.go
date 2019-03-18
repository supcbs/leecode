package main

import (
	"fmt"
)

/**

题目:
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

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

	for i:=0;i<len(nums);i++{
		tree := <-ch
		if nums[i] == -1 {
			tree.Left = nil
		} else {
			tree.Left = &TreeNode{
				Val:nums[i],
			}

			ch <- tree.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			tree.Right = nil
		} else {
			tree.Right = &TreeNode{
				Val:nums[i],
			}

			ch <- tree.Right
		}
	}

	return root
}

func main() {
	l := buildTree([]int{1})
	r := preorderTraversal2(l)
	fmt.Println(r)
}

/**
方法1:

https://blog.csdn.net/zgaoq/article/details/79089819
递归 同145、94

时间复杂度：O(n)
空间复杂度：O(1)
*/
var ret []int
func preorderTraversal(root *TreeNode) []int {
	ret := []int{}

	help(root)
	return ret
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	ret = append(ret, root.Val)
	help(root.Left)
	help(root.Right)
}


/**
方法2:

非递归。

前序遍历核心的思路是，需要一个栈，先压入root节点，后右节点，再左节点.

这里可以用数组来实现栈，终止条件是栈内没元素


时间复杂度：O(n)
空间复杂度：O(1)
*/
func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	stack = append(stack, root)

	var ret []int

	for len(stack) > 0 {
		curNode := stack[len(stack)-1]

		// 出栈
		stack = stack[:len(stack)-1]
		ret = append(ret, curNode.Val)

		if curNode.Right != nil {
			stack = append(stack, curNode.Right)
		}

		if curNode.Left != nil {
			stack = append(stack, curNode.Left)
		}
	}

	return ret
}
