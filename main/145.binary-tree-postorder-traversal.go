package main

import "fmt"

/**

题目:

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
	for i := 1; i < len(nums); i++ {
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
	l := buildTree([]int{1, 2, 3,4,5,6,7})
	r := postorderTraversal2(l)
	fmt.Println(r)
}

/**
方法1:

https://blog.csdn.net/zgaoq/article/details/79089819
递归同 94 、 144

时间复杂度：O(n)
空间复杂度：O(1)
*/
var ret []int

func postorderTraversal(root *TreeNode) []int {
	ret = []int{}
	help(root)
	return ret
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	help(root.Left)
	help(root.Right)
	ret = append(ret, root.Val)

	return
}

/**
方法2:

先将左边全部压到栈里头

然后来说取出栈节点，当当前节点的右节点是nil或者右节点等于上一个节点

则进行入ret操作，同时当前节点置为pre节点

否则的话，当前节点再放到栈里头，将当前节点置为右节点，同时将右节点的左子节点全部入栈

非递归
*/
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	cur := root
	var pre *TreeNode
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}

	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right == nil || cur.Right == pre {
			ret = append(ret, cur.Val)
			pre = cur
		} else {
			stack = append(stack, cur)
			cur = cur.Right
			for cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			}
		}
		fmt.Println(stack,ret)
	}

	return ret
}
