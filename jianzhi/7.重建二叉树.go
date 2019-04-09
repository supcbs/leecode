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
	root := &TreeNode{
		Val: nums[0],
	}
	tmp := make(chan *TreeNode, len(nums))
	tmp <- root
	for i := 1; i < len(nums); i++ {
		t := <-tmp

		if nums[i] == -1 {
			t.Left = nil
		} else {
			t.Left = &TreeNode{
				Val: nums[i],
			}
			tmp <- t.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			t.Right = nil
		} else {
			t.Right = &TreeNode{
				Val: nums[i],
			}
			tmp <- t.Right
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

/**
题目：
输入中序遍历和前序遍历的结果，构造出二叉树

分析：
先放入一个栈，然后在进行打印
*/
func main() {
	preOrder := []int{1, 2, 4, 7, 3, 5, 6, 8}
	inOrder := []int{4, 7, 2, 1, 5, 3, 8, 6}
	t := constructBinaryTree(preOrder, inOrder)
	printTree(t)
}

/**
这道题需采用递归的方式进行
1.前序遍历的第一个节点为根节点，中序遍历找到根节点后，前半部分为左子树，有半部分为右子树
*/
func constructBinaryTree(preOrder []int, inOrder []int) *TreeNode {
	if len(preOrder) != len(inOrder) || len(preOrder) <= 0 {
		return nil
	}

	var t *TreeNode
	t = &TreeNode{
		Val:preOrder[0],
	}

	i :=0
	for i=0;i<len(inOrder);i++{
		if inOrder[i] == preOrder[0] {
			break
		}
	}
	t.Left = constructBinaryTree(preOrder[1:i+1],inOrder[0:i])
	t.Right = constructBinaryTree(preOrder[i+1:],inOrder[i+1:])

	return t
}

