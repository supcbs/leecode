package main

import "fmt"

/**

题目:
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func buildTree(nums []int) *TreeNode {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	root := new(TreeNode)
//	root.Val = nums[0]
//
//	ch := make(chan *TreeNode, len(nums))
//	ch <- root
//	nums = nums[1:]
//
//	for i := 0; i < len(nums); i++ {
//		tree := <-ch
//		if nums[i] == -1 {
//			tree.Left = nil
//		} else {
//			tree.Left = &TreeNode{
//				Val: nums[i],
//			}
//
//			ch <- tree.Left
//		}
//
//		i++
//		if i == len(nums) || nums[i] == -1 {
//			tree.Right = nil
//		} else {
//			tree.Right = &TreeNode{
//				Val: nums[i],
//			}
//
//			ch <- tree.Right
//		}
//	}
//
//	return root
//}


func main() {
	t := []int{3,9,20,15,7}
	m := []int{9,3,15,20,7}
	r := buildTree(t,m)
	fmt.Println(r)
}

/**
方法1:

这道题的关键点在于

前序遍历第一个节点一定是根节点。
中序遍历找到根阶段后，的前半端一定是根节点左子树，后半段一定是右子树

时间复杂度：O(n)
空间复杂度：O(1)
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:preorder[0],
	}

	// 找出分割点
	idx := 0
	for ;preorder[0] != inorder[idx];idx++{}

	root.Left = buildTree(preorder[1:idx+1],inorder[:idx])
	root.Right = buildTree(preorder[idx+1:],inorder[idx+1:])

	return root
}
