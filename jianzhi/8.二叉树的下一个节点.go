package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Parent *TreeNode
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
				Parent: t,
			}
			tmp <- t.Left
		}

		i++
		if i == len(nums) || nums[i] == -1 {
			t.Right = nil
		} else {
			t.Right = &TreeNode{
				Val: nums[i],
				Parent: t,
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
给定一个二叉树，求其中序遍历的下一个节点。（子节点多一个指向父节点的指针）

分析：
先放入一个栈，然后在进行打印
*/
func main() {
	t := buildTree([]int{1,2,3,4,5,6,7})
	p := nextNodeInBinaryTree(t)
	fmt.Println(p.Val)
}

/**
这道题关注三种情况
1.如果当前节点有右节点，那么寻找右节点下面最左的节点，如果没有那么就是右节点了
2.如果当前没有右节点，那么可能是左节点或者最尾巴的节点。
  根据当前节点是否是父亲节点的右节点来进行区分。
  如果不是父亲节点的右节点，则父亲节点就是下一个
  如果是父亲节点的右节点，则需要往上找。制止父亲节点的右节点不等于当前节点为止
*/
func nextNodeInBinaryTree(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}

	var pNext *TreeNode
	if t.Right != nil {
		pNext = t.Right
		for pNext.Left != nil {
			pNext = pNext.Left
		}
	} else {
		parentNode := t.Parent
		curNode := t
		for parentNode != nil && parentNode.Right == curNode {
			curNode = parentNode
			parentNode = parentNode.Parent
		}

		pNext = parentNode
	}

	return pNext
}

