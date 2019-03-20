package main

import "fmt"

/**

题目:

给定一个二叉树，找出其最小深度。
最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
说明: 叶子节点是指没有子节点的节点。

示例:
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.

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
	t := buildTree([]int{1, 2})
	r := minDepth(t)
	fmt.Println(r)
}

/**
方法1:

前序遍历，然后计算最小的深度

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return 1
	}

	min := 0
	help(root, 0, &min)

	return min
}

func help(root *TreeNode, depth int, min *int) {
	if root == nil {
		return
	}

	depth += 1
	if root.Left == nil && root.Right == nil {
		if *min == 0 {
			*min = depth
		} else if depth < *min {
			*min = depth
		}
		return
	}

	help(root.Left, depth, min)
	help(root.Right, depth, min)
}
