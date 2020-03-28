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
	t := []int{1, -1, 2}
	d := buildTree(t)
	r := findMode(d)
	fmt.Println(r)
}

const MININT = ^int(^uint(0) >> 1)

func findMode(node *TreeNode) (ret []int) {
	if node == nil {
		return
	}

	var maxTimes int
	preVal := 0
	curTimes := 0
	stack := make([]*TreeNode, 0)
	cur := node
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if curTimes == 0 {
				curTimes++
				preVal = cur.Val
			} else {
				if preVal == cur.Val {
					curTimes++
				} else {
					preVal = cur.Val
					curTimes = 1
				}
			}

			if curTimes > maxTimes {
				ret = []int{cur.Val}
				maxTimes = curTimes
			} else if curTimes == maxTimes {
				ret = append(ret, cur.Val)
			}

			cur = cur.Right
			println(cur)
		}
	}
	return
}
