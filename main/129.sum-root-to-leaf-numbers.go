package main

import (
	"fmt"
)

/**

题目:
给定一个二叉树，它的每个结点都存放一个 0-9 的数字，
每条从根到叶子节点的路径都代表一个数字。
例如，从根到叶子节点路径 1->2->3 代表数字 123。
计算从根到叶子节点生成的所有数字之和。
说明: 叶子节点是指没有子节点的节点。

示例 1:
输入: [1,2,3]
    1
   / \
  2   3
输出: 25
解释:
从根到叶子节点路径 1->2 代表数字 12.
从根到叶子节点路径 1->3 代表数字 13.
因此，数字总和 = 12 + 13 = 25.

示例 2:
输入: [4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出: 1026
解释:
从根到叶子节点路径 4->9->5 代表数字 495.
从根到叶子节点路径 4->9->1 代表数字 491.
从根到叶子节点路径 4->0 代表数字 40.
因此，数字总和 = 495 + 491 + 40 = 1026.

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
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	stack = append(stack, t)
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		ret = append(ret, cur.Val)

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
	}

	fmt.Println(ret)
}

func main() {
	t := buildTree([]int{1, 2, 3})
	r := sumNumbersOK(t)
	fmt.Println(r)
}

/**
方法1:

前序遍历 + dfs 同113

时间复杂度：O(n)
空间复杂度：O(1)
*/
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var retRums []int
	var dfs func(*TreeNode, []int)
	dfs = func(node *TreeNode, nums []int) {
		if node == nil {
			return
		}

		nums = append(nums, node.Val)
		if node.Right == nil && node.Left == nil {
			num := 0
			base := 1
			for i := len(nums) - 1; i >= 0; i-- {
				num += nums[i] * base
				base *= 10
			}
			retRums = append(retRums, num)
			return
		}

		dfs(node.Left, nums)
		dfs(node.Right, nums)
	}

	dfs(root,retRums)
	ret := 0
	for i:=0;i<len(retRums);i++{
		ret+=retRums[i]
	}

	return ret
}


func sumNumbersOK(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	x := 0
	help(root.Left,root.Val,&x)
	help(root.Right,root.Val,&x)

	return x
}

func help(root *TreeNode, a int,x *int) {
	a *=10
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		fmt.Println(*x,a,root.Val)
		*x = *x + a + root.Val
		return
	}

	help(root.Left, a + root.Val, x)
	help(root.Right, a + root.Val, x)
}
