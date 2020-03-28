package main

import (
	"fmt"
	"strings"
)

/**

题目:
331. 验证二叉树的前序序列化
序列化二叉树的一种方法是使用前序遍历。当我们遇到一个非空节点时，
我们可以记录下这个节点的值。如果它是一个空节点，我们可以使用一个标记值记录，例如 #。

     _9_
    /   \
   3     2
  / \   / \
 4   1  #  6
/ \ / \   / \
# # # #   # #
例如，上面的二叉树可以被序列化为字符串 "9,3,4,#,#,1,#,#,2,#,6,#,#"，其中 # 代表一个空节点。
给定一串以逗号分隔的序列，验证它是否是正确的二叉树的前序序列化。编写一个在不重构树的条件下的可行算法。
每个以逗号分隔的字符或为一个整数或为一个表示 null 指针的 '#' 。
你可以认为输入格式总是有效的，例如它永远不会包含两个连续的逗号，比如 "1,,3" 。

示例 1:

输入: "9,3,4,#,#,1,#,#,2,#,6,#,#"
输出: true
示例 2:

输入: "1,#"
输出: false
示例 3:

输入: "9,#,#,1"
输出: false

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
	t1 := "9,3,4,#,#,1,#,#,2,#,6,#,#"
	r := isValidSerialization(t1)
	fmt.Println(r)
}

/**
方法1:

前序遍历进行判定

时间复杂度：O(n)
空间复杂度：O(h)
*/

func isValidSerialization(preorder string) bool {
	splits := strings.Split(preorder, ",")

	// 不管有没有节点，需要一个槽位
	count := 1
	for _, str := range splits {
		// 没来一个节点，都消耗一个槽位
		count--

		// 槽位如果不够了，说明不是前序遍历序列化的字符串
		if count < 0 {
			break
		}

		// 为空的节点则新增两个槽位
		if str != "#" {
			count += 2
		}
	}

	// 最后剩余槽位必须为0
	return count == 0

}
