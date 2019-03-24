package main

import "fmt"

/**

题目:
给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4

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
	t := buildTree([]int{5,1,4,-1,-1,3,6})
	r := isValidBST(t)
	fmt.Println(r)
}

/**
方法1:

中序遍历，这道题的核心在于，中序遍历是越来越大的，

时间复杂度：O(n)
空间复杂度：O(1)
*/
var pre *TreeNode
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return help(root)
}

func help(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if help(root.Left) == false {
		return false
	}

	if pre != nil && pre.Val >= root.Val {
		return false
	}

	pre = root

	return help(root.Right)
}