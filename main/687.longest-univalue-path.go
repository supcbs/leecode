package main

import "fmt"

/**
题目:
给定一个二叉树，找到最长的路径，这个路径中的每个节点具有相同值。 这条路径可以经过也可以不经过根节点。

注意：两个节点之间的路径长度由它们之间的边数表示。

示例 1:

输入:

              5
             / \
            4   5
           / \   \
          1   1   5
输出:
2
示例 2:
输入:

              1
             / \
            4   5
           / \   \
          4   4   5
输出:
2
注意: 给定的二叉树不超过10000个结点。 树的高度不超过1000。
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
	t1 := buildTree([]int{5, 4, 5, 1, 1, 5})
	r := longestUnivaluePath(t1)
	fmt.Println(r)
}

/**
方法：
核销掌握两个点：
1. 需要比较对角的时候，是否和当前最长的对比一下。
              5
             / \
            4   5
	需要比对4-5-5和累计的值哪个大
2. 返回当前两个分支中，最长的那个即可给关键节点
              5
             / \
            4   5
	比如全文返回5-5最长，为1

时间复杂度：O（n）
空间复杂度：O（h）h树高
*/
var sum int

func longestUnivaluePath(root *TreeNode) int {
	sum = 0
	help(root)

	return sum
}

func help(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := help(root.Left)
	r := help(root.Right)

	lSame := int(0)
	rSame := int(0)

	if root.Left != nil && root.Left.Val == root.Val {
		lSame = l + 1
	}

	if root.Right != nil && root.Right.Val == root.Val {
		rSame = r + 1
	}

	sum = getMax(lSame+rSame, sum)

	return getMax(lSame, rSame)
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

/**
方法同自己写的，就是他跑的比较快
---
注意一下*int当作入参怎么赋值
 */
func longestUnivaluePathFast(root *TreeNode) int {
	path := 0
	dfs(root, &path)
	return path
}

func dfs(root *TreeNode, Path *int) int{
	if root == nil{
		return 0
	}

	left := dfs(root.Left, Path)
	right := dfs(root.Right, Path)
	leftPath := 0
	rightPath := 0
	if root.Left != nil && root.Left.Val == root.Val{
		leftPath = left + 1
	}
	if root.Right != nil && root.Right.Val == root.Val{
		rightPath = right + 1
	}
	rootPath := leftPath + rightPath
	*Path = max(*Path, rootPath)
	return max(leftPath, rightPath)
}

func max(x, y int) int{
	if x > y{
		return x
	} else{
		return y
	}
}
