package main

import "fmt"

/**

题目:
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过根结点。

示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。

注意：两结点之间的路径长度是以它们之间边的数目表示。
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
	r := diameterOfBinaryTree(t)
	fmt.Println(r)
}

/**
方法1:

这道题的核心在于：
最大的值只可能存在以下情况：

1.当前的左右分支长度相加，与当前长度的长对比
          1
         / \
        2   3
   这个是为了哦去处2-1-3这种打角的情况的数值，此时就是左边加上右边即可，不需要加1！！！！因为默认都会返回1了。
   按照最简单的数数，上图情况就是1+1=2，不需要再加一个1。
2.给上级返回的时候，需要是当前值加上较大的左或者右分支数 + 1

时间复杂度：O(n)
空间复杂度：O(1)
*/
var sum int

func diameterOfBinaryTree(root *TreeNode) int {
	sum = 0
	help(root)

	return sum
}

func help(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := help(root.Left)
	r := help(root.Right)

	sum = max(l+r, sum)
	return max(l, r) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
