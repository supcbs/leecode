package main

import "fmt"

/**

题目:

给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明: 叶子节点是指没有子节点的节点。

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

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
	t := buildTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, 1})
	sum := 22
	r := pathSum(t, sum)
	fmt.Println(r)
}

/**
方法1:

判断终止条件：左右节点都为nil 并且 sum == 0

时间复杂度：O(n)
空间复杂度：O(1)
*/
func pathSum(root *TreeNode, sum int) [][]int {
	var ret [][]int
	var path []int
	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && sum == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ret = append(ret, tmp)
			return
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)

		// 循环完毕，需要将当前元素给删除掉，避免影响,下一个节点的遍历
		path = path[:len(path)-1]
	}

	dfs(root, sum)
	return ret
}
