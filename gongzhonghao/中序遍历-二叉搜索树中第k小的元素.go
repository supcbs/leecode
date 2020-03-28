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
	k := 2
	r := getKNum(d, k)
	fmt.Println(r)
}

func getKNum(node *TreeNode, k int) int {
	if node == nil {
		return 0
	}

	stack := make([]*TreeNode,0)
	cur := node
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack,cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 当k减到0的时候就是题目要的数值了
			k--
			if k == 0 {
				return cur.Val
			}
			cur = cur.Right
		}
	}
	return 0
}