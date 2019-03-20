package main

import "fmt"

/**

题目:

给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其自底向上的层次遍历为：

[
  [15,7],
  [9,20],
  [3]
]

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
	t := buildTree([]int{1, 2, 3, 4, 5})
	r := levelOrderBottom(t)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var lastNodes []*TreeNode

	lastNodes = append(lastNodes, root)

	for len(lastNodes) > 0 {
		var tmp []int
		var tmpNext []*TreeNode

		for i := 0; i < len(lastNodes); i++ {
			tmp = append(tmp, lastNodes[i].Val)

			if lastNodes[i].Left != nil {
				tmpNext = append(tmpNext, lastNodes[i].Left)
			}
			if lastNodes[i].Right != nil {
				tmpNext = append(tmpNext, lastNodes[i].Right)
			}
		}

		ret = append(ret, tmp)
		lastNodes = tmpNext
	}

	fmt.Println(ret)
	var r [][]int
	for i := len(ret) - 1; i >= 0; i-- {
		r = append(r, ret[i])
	}

	return r
}
