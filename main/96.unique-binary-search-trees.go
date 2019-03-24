package main

import "fmt"

/**

题目:

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:
输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

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
	t := 3
	r := numTrees(t)
	fmt.Println(r)
}

/**
方法1:

g(n) = g(0)*g(n-1) + g(1)*g(n-2) + +g(n)*g(0)

时间复杂度：O(n)
空间复杂度：O(1)
*/
func numTrees(n int) int {
	ret := 0
	if n == 0 {
		return 1
	} else if n <= 2{
		return n
	}

	for i:=0;i<n;i++{
		ret += numTrees(i)*numTrees(n-i-1)
	}

	return ret
}

func numTreesOK(n int) int {
	G := make([]int, n+1)
	G[0],G[1] = 1,1
	for i:=2;i<n;i++{
		for j:=0;j<i;j++{
			G[i] += G[j]*G[i-j-1]
		}
	}

	return G[n]
}
