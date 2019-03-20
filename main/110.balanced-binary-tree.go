package main

import (
	"fmt"
)

/**

题目:
给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:
给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

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
	r := isBalanced(t)
	fmt.Println(r)
}

/**
方法1:

1.采用后序遍历的方式
2.然后比较每一层差值。
3.返回值是当前差值比较大的+1返回

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return false
	}

	ret := true
	help(root, &ret)

	return ret
}

func help(root *TreeNode, ret *bool) int {
	if *ret == false {
		return 0
	}

	if root == nil {
		return 0
	}

	left := help(root.Left, ret)
	right := help(root.Right, ret)

	if left-right > 1 || right-left > 1 {
		*ret = false
		return 0
	}

	if left > right {
		return left + 1

	}

	return right + 1
}

func isBalancedOK(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if left-right <= 1 && left-right >= -1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
	return false
}

func depth(n *TreeNode) int {
	if n == nil {
		return 0
	} else {
		left := depth(n.Left)
		right := depth(n.Right)
		if left >= right {
			return left + 1
		} else {
			return right + 1
		}
	}
}
