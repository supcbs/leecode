package main

import "fmt"

/**

题目:

给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。

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
	t := buildTree([]int{3,9,20,-1,-1,1,1})
	r := maxDepth(t)
	fmt.Println(r)
}

/**
方法1:

其实也是前序遍历
同111

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	help(root, 0, &max)

	return max
}

func help(root *TreeNode, depth int, max *int) {
	if root == nil {
		return
	}

	depth++
	if root.Right == nil && root.Left == nil {
		if depth > *max {
			*max = depth
		}
		return
	}
	help(root.Left, depth,max)
	help(root.Right, depth,max)
}


func maxDepthOK(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}


func max(left, right int) int {
	if left > right {
		return left
	}
	return right
}


