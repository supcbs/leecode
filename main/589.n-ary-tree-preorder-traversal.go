package main

import "fmt"

/**

给定一个 N 叉树，返回其节点值的前序遍历。

例如，给定一个 3叉树 :

示例 1:
输入:       1
          / | \
         3  2  4
        / \
       5   6

返回其前序遍历: [1,3,5,6,2,4]。

说明: 递归法很简单，你可以使用迭代法完成此题吗?
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

type Node struct {
	Val      int
	Children []*Node
}

func main() {
	//t1 := buildTree([]int{1, 2, 1})
	//r := preorder(t1)
	//fmt.Println(r)
}

/**
方法1:

前序遍历进行判定

时间复杂度：O(n)
空间复杂度：O(h)
*/

func preorder(root *Node) []int {
	res := make([]int,0)
	help(root, &res)
	return res
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
func help(root *Node, res *[]int) {
	if root == nil {
		return
	}

	// 添加值到返回值当中
	*res = append(*res, root.Val)

	// 多个节点，顺序append即可，
	for _,v := range root.Children {
		help(v, res)
	}
}

func preorderOK(root *Node) []int {
	res := make([]int,0)
	if root == nil{
		return res
	}
	res = append(res,root.Val)
	n :=len(root.Children)
	for i:=0;i<n;i++{
		res = append(res,preorder(root.Children[i])...)
	}
	return res
}