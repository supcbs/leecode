package main

import (
	"fmt"
)

/**

题目:
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/

type ListNode struct {
	Val int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val:nums[0],
	}
	tmp := root
	for i:=1;i<len(nums);i++{
		tmp.Next = &ListNode{
			Val:nums[i],
		}
		tmp = tmp.Next
	}

	return root
}

func main() {
	l1:=buildList([]int{1,2,4})
	l2:=buildList([]int{1,3,4})
	r := mergeTwoLists(l1,l2)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	h := &ListNode{}
	p := h


	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			p = p.Next
			l1 = l1.Next
		} else {
			p.Next = l2
			p = p.Next
			l2 = l2.Next
		}
	}

	for l1 != nil {
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}

	return h.Next
}

/**
递归
---
有点动态规划的感觉
当的节点的下一个一定是两个链表中比较小的一个
 */
func mergeTwoListsOK(l1 *ListNode, l2 *ListNode) *ListNode {
	// 说明另一个更长
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		return &ListNode{l1.Val,mergeTwoListsOK(l1.Next,l2)}
	} else {
		return &ListNode{l2.Val,mergeTwoListsOK(l1,l2.Next)}
	}
}