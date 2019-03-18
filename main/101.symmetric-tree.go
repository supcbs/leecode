package main

import "fmt"

/**

题目:

给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

说明:
如果你可以运用递归和迭代两种方法解决这个问题，会很加分。

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
	t := buildTree([]int{1, 2, 2, 3, 4, 4, 3})
	//t = buildTree([]int{1,2,2,-1,3,-1,3})
	r := isSymmetricOK(t)
	fmt.Println(r)
}

/**
方法1:

广度遍历，然后进行入栈出栈操作
---
这种姿势不优雅，看递归的方式

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var lastArr []*TreeNode
	lastArr = append(lastArr, root)

	for len(lastArr) > 0 {
		var tmpArr []*TreeNode
		for i := 0; i <= len(lastArr)/2; i++ {
			if lastArr[i].Val != lastArr[len(lastArr)-i-1].Val {
				return false
			}
		}

		flag := false
		for i := 0; i < len(lastArr); i++ {
			if lastArr[i].Left == nil {
				tmpArr = append(tmpArr, &TreeNode{Val: -1})
			} else {
				tmpArr = append(tmpArr, lastArr[i].Left)
				flag = true
			}

			if lastArr[i].Right == nil {
				tmpArr = append(tmpArr, &TreeNode{Val: -1})
			} else {
				tmpArr = append(tmpArr, lastArr[i].Right)
				flag = true
			}
		}

		if flag == true {
			lastArr = tmpArr
		} else {
			lastArr = nil
		}
	}

	return true
}

/**
方法2:

对于递归来说：
其实就是相当于比对两个左节点和右节点是否相等

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isSymmetricOK(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	return help(root.Left, root.Right)
}

func help(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 != nil && t2 != nil && t1.Val == t2.Val {
		return help(t1.Left, t2.Right) && help(t2.Left, t1.Right)
	}

	return false
}
