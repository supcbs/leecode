package main

import "fmt"

/**

题目:
给给定一个所有节点为非负值的二叉搜索树，求树中任意两节点的差的绝对值的最小值。

示例 :

输入:

   1
    \
     3
    /
   2

输出:
1

解释:
最小绝对差为1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
注意: 树中至少有2个节点。

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
	t := buildTree([]int{5,1,4})
	r := getMinimumDifference(t)
	fmt.Println(r)
}

/**
方法：

这道题一开始想歪了，实际上由于是搜索二叉树，所以任意节点差最小值一定出现在相邻的两个节点中。

中序遍历 + 记录上一个节点的值做差值即可

时间复杂度：O(n)
空间复杂度：O(1)
 */
var prev int
var res int
func getMinimumDifference(root *TreeNode) int {
	prev = -1
	res = int(^uint(0) >> 1)
	help(root)
	return res
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	help(root.Left)

	// 这一步是关键
	if prev >= 0 {
		if root.Val - prev < res {
			res = root.Val - prev
		}
	}
	prev = root.Val
	help(root.Right)
	return
}






