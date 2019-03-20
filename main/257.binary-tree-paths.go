package main

import (
	"fmt"
	"strconv"
)

/**

题目:
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。
示例:
输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]
解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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
	r := binaryTreePaths(t)
	fmt.Println(r)
}

/**
方法1:

到达了根节点，遇到节点进行返回，然后前头插入法

时间复杂度：O(n)
空间复杂度：O(1)
*/
func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	if root == nil {
		return ret
	}

	// 到达了叶子结点
	str := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{str}
	}

	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)

	fmt.Println(left, right,ret)
	str = str + "->"
	for i := 0; i < len(left); i++ {
		ret = append(ret, str+left[i])
	}
	for i := 0; i < len(right); i++ {
		ret = append(ret, str+right[i])
	}

	fmt.Println("===",left, right,ret)
	return ret
}

/**
方法1:

核心思路是到达根节点的时候，进行添加路径

时间复杂度：O(n)
空间复杂度：O(1)
*/

var ret []string
func binaryTreePathsOK(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	help(root, "")
	return ret
}

func help(root *TreeNode, path string) {
	if root == nil {
		return
	}

	// 拼接字符串
	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path = path + "->" + strconv.Itoa(root.Val)
	}

	// 到达终止条件则添加到ret中
	if root.Left == nil && root.Right == nil {
		ret = append(ret, path)
		return
	}

	help(root.Left, path)
	help(root.Right, path)
}
