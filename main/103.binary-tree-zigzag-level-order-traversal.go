package main

import "fmt"

/**

题目:

给定一个二叉树，返回其节点值的锯齿形层次遍历。
（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层次遍历如下：

[
  [3],
  [20,9],
  [15,7]
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
	t := buildTree([]int{3, 9, 20, -1, -1, 15, 7})
	t = buildTree([]int{1,2,3,4, -1, -1, 5})
	r := zigzagLevelOrder(t)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ret [][]int
	var lastNodes []*TreeNode

	lastNodes = append(lastNodes, root)
	zheng := true
	for len(lastNodes) > 0 {
		var tmp []int
		var tmpNext []*TreeNode

		fmt.Println(lastNodes)
		if zheng {
			for i := 0; i < len(lastNodes); i++ {
				fmt.Println(i, lastNodes)
				tmp = append(tmp, lastNodes[i].Val)

				if lastNodes[i].Left != nil {
					tmpNext = append(tmpNext, lastNodes[i].Left)
				}
				if lastNodes[i].Right != nil {
					tmpNext = append(tmpNext, lastNodes[i].Right)
				}
			}
			zheng = false
		} else {
			for i := len(lastNodes) - 1; i >= 0; i-- {
				tmp = append(tmp, lastNodes[i].Val)
				if lastNodes[i].Right != nil {
					tmpNext = append([]*TreeNode{lastNodes[i].Right}, tmpNext[:]...)
				}
				if lastNodes[i].Left != nil {
					tmpNext = append([]*TreeNode{lastNodes[i].Left}, tmpNext[:]...)
				}
			}
			zheng = true
		}

		lastNodes = tmpNext
		ret = append(ret, tmp)
	}

	return ret
}


func zigzagLevelOrderOK(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}

	var res [][]int
	var dfs func(*TreeNode,int)

	dfs = func(node *TreeNode, level int) {
		if node == nil{
			return
		}

		if level>=len(res){
			res = append(res,[]int{})
		}

		if level % 2 == 0{
			res[level] = append(res[level],node.Val)
		}else {
			tmp := make([]int,len(res[level])+1)
			tmp[0] = node.Val
			copy(tmp[1:],res[level])
			res[level] = tmp
		}

		dfs(node.Left,level+1)
		dfs(node.Right,level+1)
	}
	dfs(root,0)
	return res
}
