package main

import "fmt"

/**

题目:
二叉搜索树中的两个节点被错误地交换。

请在不改变其结构的情况下，恢复这棵树。

示例 1:

输入: [1,3,null,null,2]

   1
  /
 3
  \
   2

输出: [3,1,null,null,2]

   3
  /
 1
  \
   2
示例 2:

输入: [3,1,4,null,null,2]

  3
 / \
1   4
   /
  2

输出: [2,1,4,null,null,3]

  2
 / \
1   4
   /
  3
进阶:

使用 O(n) 空间复杂度的解法很容易实现。
你能想出一个只使用常数空间的解决方案吗？

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
	t := buildTree([]int{1, 2, 3, 4, 5})
	r := recoverTree
	fmt.Println(r)
}

/**
方法1:

这道题的核心在于，反正和二叉搜索树相关的大概率是中序遍历

由于中序遍历一定是递增的，当出现前一个大于下一个的时候，所以需要记住前一个节点的数，然后即为交换的目标1，
第二个数则是后面最后一个出现的，如果后续还有出现倒叙的情况，则更新目标2

时间复杂度：O(n)
空间复杂度：O(1)
*/
var s1, s2, pre *TreeNode

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

	s1, s2, pre = nil, nil, nil
	help(root)
	tmp := s1.Val
	s1.Val = s2.Val
	s2.Val = tmp
	return
}

func help(root *TreeNode) {
	if root == nil {
		return
	}

	help(root.Left)
	if pre != nil && pre.Val > root.Val {
		if s1 == nil {
			s1 = pre
			s2 = root
		} else {
			s2 = root
		}
	}
	pre = root
	help(root.Right)
}
