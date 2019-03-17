package main

import (
	"fmt"
)

/**

题目:

*/

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l:=buildList([]int{1,2,3,3,4,4,5})
	r := deleteDuplicates(l)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
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

/**
方法1:

这道题的关键在于：上一指针在何时发生移位。

移位动作 => prev = prev.Next
只有当当前位 == prev.Next的时候，说明cur指向的数字是唯一的


时间复杂度：O(n)
空间复杂度：O(1)
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var h ListNode
	h.Next = head
	prev := &h
	cur := prev.Next
	for cur != nil {
		for cur.Next != nil && cur.Next.Val == cur.Val {
			cur = cur.Next
		}

		if prev.Next == cur {
			prev = prev.Next
		} else {
			prev.Next = cur.Next
		}

		cur = cur.Next
	}

	return h.Next
}
