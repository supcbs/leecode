package main

import (
	"fmt"
)

/**

题目:
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

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
	t := buildTree([]int{1, 2, 3, 4, 5})
	r := temp1(t)
	fmt.Println(r)
}

/**
方法1:

同108，关键点在于先将链表转为数组

时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	root := help(nums)

	return root
}

func help(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2
	root := &TreeNode{
		Val: nums[mid],
	}

	if mid > 0 {
		root.Left = help(nums[0:mid])
	}
	if mid < len(nums)-1 {
		root.Left = help(nums[mid+1:])
	}

	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{
			Val: head.Val,
			Left: nil,
			Right:  nil,
		}
	}

	f := head
	s := head
	c := s
	for f.Next != nil {
		f = f.Next
		if f.Next != nil {
			f = f.Next
			c = s
			s = s.Next
		}
	}

	nextHead := s.Next
	// 这一步极其关键，代表去掉这个s的元素
	c.Next = nil
	root := &TreeNode{
		Val: s.Val,
	}

	if head != s {
		root.Left = sortedListToBST(head)
	}
	root.Right = sortedListToBST(nextHead)
	return root
}
