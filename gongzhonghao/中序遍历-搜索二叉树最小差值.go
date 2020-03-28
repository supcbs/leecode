package main

import "fmt"

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
	t := []int{1,-1,3,2}
	d := buildTree(t)
	r := getMinimumDifference(d)
	fmt.Println(r)
}

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