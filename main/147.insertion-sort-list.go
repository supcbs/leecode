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
	l:=buildList([]int{3,4,2,1,5})
	r := insertionSortListOK(l)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

首先插入排序的核心是从第二个位置开始往前找，找到合适的位置进行放置

将第一个元素放到单独的链里头，然后从第二个元素开始遍历这个链进行插入

时间复杂度：O(?)
空间复杂度：O(1)
*/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 第一个元素提取出来
	newL := &ListNode{}
	p:=newL
	p.Next = head
	p = p.Next

	head = head.Next
	p.Next = nil

	for head != nil {
		p := newL.Next
		prev := newL

		for p != nil && head.Val > p.Val {
			prev = p
			p = p.Next
		}

		tmp := head.Next
		head.Next = p
		prev.Next = head
		head = tmp
	}

	return newL.Next
}


/**
方法1:

首先插入排序的核心是从第二个位置开始往前找，找到合适的位置进行放置

先将链表进行存住，然后从第一个开始往后面进行比较大小，
当后面的小于前面的时候，将这个节点拿出（前后进行闭合）
然后那这个节点进行通头开始的比对，放到合适的位置

时间复杂度：O(?)
空间复杂度：O(1)
*/
func insertionSortListOK(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 第一个元素提取出来
	newL := &ListNode{
		Next:head,
	}

	cur := head
	for cur != nil && cur.Next!=nil{
		p := cur.Next
		if p.Val > cur.Val {
			cur = cur.Next
			p = p.Next
			continue
		}

		cur.Next = p.Next

		pre := newL
		next := newL.Next
		for p.Val > next.Val {
			pre = pre.Next
			next = next.Next
		}

		pre.Next = p
		p.Next = next
	}

	return newL.Next
}
