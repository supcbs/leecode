package main

import "fmt"

/**
题目:
给出二叉树的根节点 root，树上每个节点都有一个不同的值。
如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
返回森林中的每棵树。你可以按任意顺序组织答案。

示例：
输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
输出：[[1,2,null,4],[6],[7]]


提示：

树中的节点数最大为 1000。
每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
to_delete.length <= 1000
to_delete 包含一些从 1 到 1000、各不相同的值。
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
	t1 := buildTree([]int{1, 2, 3, 4, 5, 6, 7})
	toDelete := []int{3, 5}
	r := delNodes(t1, toDelete)
	fmt.Println(r[0],r[1])
}

/**
方法：
这道题的关键在于，能否想到用递归的方式来进行数据的采集。

1. 存在删除节点中，则返回nil，并且把左、右节点放到目标节点中（需要不为nil）
2. 核心采用的是后序遍历，因为需要先分析子树，才能端到你

*/
var nodes []*TreeNode
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	nodes = make([]*TreeNode, 0)
	node := proc(root, to_delete)

	if node != nil {
		nodes = append(nodes, node)
	}

	return nodes
}

func proc(root *TreeNode, to_delete []int) *TreeNode {
	if root == nil {
		return nil
	}

	left := proc(root.Left, to_delete)
	right := proc(root.Right, to_delete)

	if inArr(root.Val, to_delete) {
		if left != nil {
			nodes = append(nodes, left)
		}

		if right != nil {
			nodes = append(nodes, right)
		}

		root = nil
	}

	return root
}

func inArr(num int, toDelete []int) bool {
	for _, v := range toDelete {
		if v == num {
			return true
		}
	}

	return false
}
