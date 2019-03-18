package main

import "fmt"

/**

题目:
给定一个二叉树，返回它的中序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]

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
	l := buildTree([]int{1, 2, 3})
	r := inorderTraversal2(l)
	fmt.Println(r)
}

/**
方法1:

https://blog.csdn.net/zgaoq/article/details/79089819
递归同144题

时间复杂度：O(n)
空间复杂度：O(1)
*/
var ret []int

func inorderTraversal(root *TreeNode) []int {
	ret = []int{}
	help(root)
	return ret
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	help(root.Left)
	ret = append(ret, root.Val)
	help(root.Right)
}

/**
方法2:

非递归：
中序遍历的原则是左边的节点优先，不存在左边节点了，则取出栈里头元素指向右边节点

时间复杂度：O(n)
空间复杂度：O(1)
*/

func inorderTraversal2(root *TreeNode) []int {
	ret = []int{}

	var stack []*TreeNode

	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = append(ret, cur.Val)
			cur = cur.Right
		}
	}

	return ret
}
