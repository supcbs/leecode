package main

import "fmt"

/**

题目:
给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，
这条路径上所有节点值相加等于目标和。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \      \
        7    2      1

返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。
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
	t := buildTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, 1})
	sum := 22
	t = buildTree([]int{1,2})
	sum = 1
	r := hasPathSum(t, sum)
	fmt.Println(r)
}

/**
方法1:

前序遍历，进行数组累加，发现等于目标数就返回
注意题目是一定要到跟！！！

时间复杂度：O(n)
空间复杂度：O(1)
*/
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}

	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
}

func hasPathSumOK(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	if root.Left == nil && root.Right == nil{
		return root.Val == sum
	}
	sum -= root.Val
	fmt.Println(sum)
	flag := hasPathSum(root.Left,sum)
	if flag == false {
		flag = hasPathSum(root.Right,sum)
	}
	return flag
}
