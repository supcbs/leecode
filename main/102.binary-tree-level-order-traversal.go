package main

import (
	"fmt"
)

/**

题目:

给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
例如:
给定二叉树: [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7

返回其层次遍历结果：
[
  [3],
  [9,20],
  [15,7]
]

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = nums[0]
	ch := make(chan *TreeNode, len(nums))
	ch <- root

	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		cur := <-ch

		if nums[i] == -1 {
			cur.Left = nil
		} else {
			cur.Left = &TreeNode{
				Val: nums[i],
			}
			ch <- cur.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			cur.Right = nil
		} else {
			cur.Right = &TreeNode{
				Val: nums[i],
			}
			ch <- cur.Right
		}
	}

	return root
}

func main() {
	t := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
	r := levelOrderOK(t)
	fmt.Println(r)
}

/**
方法1:

遍历的时候，关键是记住上一次存进去的节点信息。
截至条件上存进去的节点信息为空

时间复杂度：O(n)
空间复杂度：O(1)
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var lastNodes []*TreeNode

	lastNodes = append(lastNodes, root)
	for len(lastNodes) > 0 {
		var tmpNext []*TreeNode
		var tmpRet []int
		for i := 0; i < len(lastNodes); i++ {
			tmpRet = append(tmpRet, lastNodes[i].Val)

			if lastNodes[i].Left != nil {
				tmpNext = append(tmpNext, lastNodes[i].Left)
			}

			if lastNodes[i].Right != nil {
				tmpNext = append(tmpNext, lastNodes[i].Right)
			}
		}

		lastNodes = tmpNext
		ret = append(ret, tmpRet)
	}

	return ret
}

/**
这道题是典型的广度遍历 BFS
 */
func levelOrderOK(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	type BFS struct{
		node *TreeNode
		depth int
	}

	q := []BFS{
		{node:root,depth:0},
	}

	var ret [][]int
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		node := p.node

		if p.depth == len(ret) {
			ret = append(ret, []int{node.Val})
		} else {
			ret[p.depth] = append(ret[p.depth], node.Val)
		}

		if node.Left != nil {
			q = append(q,BFS{
				node:node.Left,
				depth:p.depth+1,
			})
		}

		if node.Right != nil {
			q = append(q,BFS{
				node:node.Right,
				depth:p.depth+1,
			})
		}
	}

	return ret
}
