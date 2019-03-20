package main

import "fmt"

/**

题目:
给定一个非空二叉树，返回其最大路径和。
本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。
该路径至少包含一个节点，且不一定经过根节点。

示例 1:
输入: [1,2,3]

       1
      / \
     2   3

输出: 6

示例 2:
输入: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15   7

输出: 42

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
	t := buildTree([]int{-10, 9, 20, -1, -1, 15, 7})
	t = buildTree([]int{-2, 1})
	r := maxPathSum(t)
	fmt.Println(r)
}

/**
方法1:

这道题的核心在于：
最大的值只可能存在以下情况：

1.当前的值加上左分支，或者当前值加上右分支，当前值加上左右分支
2.给上级返回的时候，需要是当前值加上较大的左或者右分支，或者是当前值

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ret := root.Val
	help(root, &ret)
	return ret
}

func help(root *TreeNode, ret *int) int {
	if root == nil {
		return 0
	}

	left := max(0, help(root.Left, ret))
	right := max(0, help(root.Right, ret))
	*ret = max(*ret, root.Val+left+right)

	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
