package main

import "fmt"

/**

题目:
我们可以为二叉树 T 定义一个翻转操作，如下所示：选择任意节点，然后交换它的左子树和右子树。
只要经过一定次数的翻转操作后，能使 X 等于 Y，我们就称二叉树 X 翻转等价于二叉树 Y。
编写一个判断两个二叉树是否是翻转等价的函数。这些树由根节点 root1 和 root2 给出。

示例：

输入：root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
输出：true
解释：We flipped at nodes with values 1, 3, and 5.

 

提示：
每棵树最多有 100 个节点。
每棵树中的每个值都是唯一的、在 [0, 99] 范围内的整数。
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
	t1 := buildTree([]int{1, 2, 2, 3, 4, 4, 3})
	t2 := buildTree([]int{1,2,2,-1,3,-1,3})
	r := flipEquiv(t1,t2)
	fmt.Println(r)
}

/**
方法1:

这道题中的在于，如果翻转n次可以相等的话，那么n可以是0-n次

那只有两中情况，要不翻转、要不不翻转
1.不翻转的时候，两个左边肯定相等、或者两个右边肯定相等
1.翻转的时候，两个左右肯定对调相等

边界条件：
1. 两个都为nil则ok,说明到顶了
2. 其一为空、或者值不等，则no

时间复杂度：O(n)
空间复杂度：O(h)
*/
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}

	l1 := flipEquiv(root1.Left, root2.Left)
	l2 := flipEquiv(root1.Left, root2.Right)
	r1 := flipEquiv(root1.Right, root2.Right)
	r2 := flipEquiv(root1.Right, root2.Left)

	return ((l1 && r1) || (l2 && r2)) && (root1.Val == root2.Val)
}