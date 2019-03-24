package main

import "fmt"

/**

题目:
给定一个整数 n，生成所有由 1 ... n 为节点所组成的二叉搜索树。

示例:

输入: 3
输出:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
解释:
以上的输出对应以下 5 种不同结构的二叉搜索树：

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

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
	t := 3
	r := generateTrees(t)
	fmt.Println(r)
}

/**
方法1:

这道题，需要进行后续遍历

其核心思想是由1开始到n，当左边是1 -> i-1的时候，右边是i+1 -> max

时间复杂度：O(n)
空间复杂度：O(1)
*/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return help(1, n)
}

func help(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var res []*TreeNode
	for i := start; i <= end; i++ {
		left := help(start, i-1)
		right := help(i+1, end)

		for j := 0; j < len(left); j++ {
			for k := 0; k < len(left); k++ {
				root := &TreeNode{
					Val:   i,
					Left:  left[j],
					Right: right[k],
				}

				res = append(res, root)
			}
		}
	}

	return res
}
