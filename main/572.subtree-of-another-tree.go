package main

import "fmt"

/**

题目:

给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
给定的树 s:

     3
    / \
   4   5
  / \
 1   2
给定的树 t：

   4
  / \
 1   2
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 2:
给定的树 s：

     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：

   4
  / \
 1   2
返回 false。

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
	t1 := buildTree([]int{1, 2, 1})
	t2 := buildTree([]int{1, 1, 2})
	r := isSubtree(t1, t2)
	fmt.Println(r)
}

/**
方法1:

前序遍历进行判定

时间复杂度：O(n)
空间复杂度：O(h)
*/
func isSubtree(s *TreeNode, t *TreeNode) bool {
	// 目标节点为nil，则直接就是true
	if t == nil {
		return true
	}

	// 源节点为nil，直接就是false
	if s == nil {
		return false
	}

	// 判断当前节点是否一样 ||  左节点与目标节点比 || 右节点与目标节点比
	return isSubtreeHelper(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}


func isSubtreeHelper(s, p *TreeNode) bool {
	if p == nil && s == nil {
		return true
	}
	if p == nil || s == nil {
		return false
	}

	// 值必须相等，双递归剩下的左右节点也必须相等
	return s.Val == p.Val && isSubtreeHelper(s.Left, p.Left) && isSubtreeHelper(s.Right, p.Right)
}
