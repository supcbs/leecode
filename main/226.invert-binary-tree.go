package main

import "fmt"

/**

题目:
翻转一棵二叉树。

示例：
输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9

输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：
谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。
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
	r := invertTree(t)
	printTree(r)
}

/**
方法1:

翻转很简单其实就是递归的时候，当前节点的左右节点互换

1.当前节点的左右节点互换
2.进行左节点的互换
3.进行右节点的缓缓

时间复杂度：O(n)
空间复杂度：O(1)
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	help(root)
	//root.Left,root.Right = root.Right,root.Left
	//
	//root.Left = invertTree(root.Left)
	//root.Right = invertTree(root.Right)
	return root
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	root.Left,root.Right = root.Right,root.Left

	help(root.Left)
	help(root.Right)
}

