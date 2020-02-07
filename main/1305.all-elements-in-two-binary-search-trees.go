package main

import "fmt"

/**
题目:
给你 root1 和 root2 这两棵二叉搜索树。

请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。

示例 1：

    2     1
   1 4   0 3
输入：root1 = [2,1,4], root2 = [1,0,3]
输出：[0,1,1,2,3,4]
示例 2：

输入：root1 = [0,-10,10], root2 = [5,1,7,0,2]
输出：[-10,0,0,1,2,5,7,10]
示例 3：

输入：root1 = [], root2 = [5,1,7,0,2]
输出：[0,1,2,5,7]
示例 4：

输入：root1 = [0,-10,10], root2 = []
输出：[-10,0,10]
示例 5：

输入：root1 = [1,null,8], root2 = [8,1]
输出：[1,1,8,8]

提示:
每棵树最多有 5000 个节点。
每个节点的值在 [-10^5, 10^5] 之间。
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
	t1 := buildTree([]int{2,1,4})
	t2 := buildTree([]int{1,0,3})
	r := getAllElements(t1,t2)
	fmt.Println(r)
}

/**
方法1：
最简单的就是：两遍前序遍历+归并排序

 */

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	stack1 := make([]int,0)
	stack2 := make([]int,0)
	inOrder(root1,&stack1)
	inOrder(root2,&stack2)

	ret := make([]int,0)
	var index1,index2 int
	for len(ret) != len(stack1) + len(stack2) {
		if len(stack1) > index1 && len(stack2) > index2 {
			if stack1[index1] <=  stack2[index2] {
				ret = append(ret, stack1[index1])
				index1++
			} else {
				ret = append(ret, stack2[index2])
				index2++
			}
		}

		if len(stack1) > index1{
			ret = append(ret, stack1[index1])
			index1++
		}

		if  len(stack2) > index2{
			ret = append(ret, stack2[index2])
			index2++
		}
	}

	return ret
}

func inOrder(node *TreeNode,ret *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left,ret)
	*ret = append(*ret, node.Val)
	inOrder(node.Right, ret)
	return
}

/**
方法2：
这里写一个最佳的：总体思路还是前序遍历+归并排序，但是是归并排序插在前序遍历中间

 */
func getAllElementsNo(root1 *TreeNode, root2 *TreeNode) []int {
	stack1 := smallest(root1,[]*TreeNode{})
	stack2 := smallest(root2,[]*TreeNode{})

	var ret []int
	for len(stack1) > 0 || len(stack2) > 0 {
		var num1,num2 int
		if len(stack1) > 0 {
			num1 = stack1[len(stack1)-1].Val
		} else {
			num1 = 0
		}

		if len(stack2) > 0 {
			num2 = stack2[len(stack2)-1].Val
		} else {
			num2 = 0
		}

		if num1 >= num2 {
			n := stack2[len(stack2)-1]
			ret = append(ret, num2)
			stack2 = stack2[0:len(stack2)-1]
			smallest(n.Right, stack2)
		} else {
			n := stack1[len(stack1)-1]
			ret = append(ret, num1)
			stack1 = stack1[0:len(stack1)-1]
			smallest(n.Right, stack1)
		}
	}

	return ret
}

func smallest(node *TreeNode, ret []*TreeNode) []*TreeNode {
	if node == nil {
		return ret
	}

	ret = append(ret,node)
	smallest(node.Left, ret)

	return ret
}
