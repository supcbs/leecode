package main

import "fmt"

/**

题目:

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
	t := []int{9, 3, 15, 20, 7}
	t2 := []int{9, 15, 7, 20, 3}
	r := buildTree(t, t2)
	fmt.Println(r)
}

/**
方法1:

和105很像

后续遍历的最后一个一定是根节点
中序遍历查找分割点，前半段是左子树，右半段是右子树

时间复杂度：O(n)
空间复杂度：O(1)
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:postorder[len(postorder)-1],
	}

	// 寻找分割点
	idx := 0
	for ;postorder[len(postorder)-1]!=inorder[idx];idx++{}

	root.Left = buildTree(inorder[:idx],postorder[:idx])
	root.Right = buildTree(inorder[idx+1:],postorder[idx:len(postorder)-1])
	return root
}
