package main

import "fmt"

/**

题目:
给出一个完全二叉树，求出该树的节点个数。

说明：

完全二叉树的定义如下：在完全二叉树中，除了最底层节点可能没填满外，
其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
若最底层为第 h 层，则该层包含 1~ 2h 个节点。

示例:

输入:
    1
   / \
  2   3
 / \  /
4  5 6

输出: 6


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
	t := buildTree([]int{1, 2, 3, 4, 5})
	r := countNodes(t)
	fmt.Println(r)
}

/**
方法1:

节点数量，这个直接使用前序遍历即可

时间复杂度：O(n)
空间复杂度：O(1)
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := countNodes(root.Left)
	rightCount := countNodes(root.Right)

	return leftCount + rightCount+1
}
