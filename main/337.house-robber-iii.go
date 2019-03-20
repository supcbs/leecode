package main

import "fmt"

/**

题目:
在上次打劫完一条街道之后和一圈房屋后，小偷又发现了一个新的可行窃的地区。
这个地区只有一个入口，我们称之为“根”。 除了“根”之外，
每栋房子有且只有一个“父“房子与之相连。一番侦察之后，
聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。
如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。
计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

示例 1:
输入: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

输出: 7
解释: 小偷一晚能够盗取的最高金额 = 3 + 3 + 1 = 7.

示例 2:
输入: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

输出: 9
解释: 小偷一晚能够盗取的最高金额 = 4 + 5 = 9.

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
	t := buildTree([]int{1, 2, 3, 4, 5})
	r := rob(t)
	fmt.Println(r)
}

/**
方法1:

能抢到的最大金额为二者其一

1.该节点 + 左孩子的左右节点之和 + 右孩子的左右节点之和
2.其左右孩子之和


对于一个以 node 为根节点的二叉树而言，如果尝试偷取 node 节点，
那么势必不能偷取其左右子节点，然后继续尝试偷取其左右子节点的左右子节点。
如果不偷取该节点，那么只能尝试偷取其左右子节点。
比较两种方式的结果，谁大取谁。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0
	if root.Left != nil {
		sum = sum + rob(root.Left.Left) + rob(root.Left.Right)
	}

	if root.Right != nil {
		sum = sum + rob(root.Right.Left) + rob(root.Right.Right)
	}

	sum += root.Val
	t := rob(root.Left) + rob(root.Right)
	if sum > t {
		return sum
	}

	return t
}


func robOK(root *TreeNode) int {
	if root == nil {
		return 0
	}

	robRoot, notRobRoot := choose(root)

	if robRoot > notRobRoot {
		return robRoot
	} else {
		return notRobRoot
	}
}

func choose(node *TreeNode) (rob int, notRob int) {
	if node == nil {
		return 0, 0
	}

	robL, notRobL := choose(node.Left)
	robR, notRobR := choose(node.Right)

	rob, notRob = node.Val, 0
	rob += notRobL + notRobR

	if robL > notRobL {
		notRob += robL
	} else {
		notRob += notRobL
	}

	if robR > notRobR {
		notRob += robR
	} else {
		notRob += notRobR
	}

	return
}
