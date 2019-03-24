package main

import "fmt"

/**

题目:
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5

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
	t := []int{-10, -3, 0, 5, 9}
	r := sortedArrayToBST(t)
	fmt.Println(r)
}

/**
方法1:

关键核心在于两点：
1.需要是平衡二叉树
2.数组已经是排好序的了

所以只需要：
1.每次取中间，将左右部分的中间值赋值为当前节点的左右值即可

时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortedArrayToBST2(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}

	root := help(nums, 0, len(nums)-1)
	return root
}

func help(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	fmt.Println(mid)
	root := &TreeNode{
		Val: nums[mid],
	}

	root.Left = help(nums, left, mid-1)
	root.Right = help(nums, mid+1, right)

	return root
}


func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val:nums[0]}
	}

	mid := len(nums)/2
	root := &TreeNode{
		Val:nums[mid],
	}

	if mid > 0 {
		root.Left = sortedArrayToBST(nums[0:mid])
	}

	if mid < len(nums)-1 {
		root.Right = sortedArrayToBST(nums[mid+1:])
	}

	return root
}
