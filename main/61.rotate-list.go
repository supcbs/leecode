package main

import "fmt"

/**

题目:
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

示例 2:
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

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
	//l:=buildList([]int{1,2,3,4,5})
	//r := rotateRight(l,2)
	l:=buildList([]int{0,1,2})
	r := rotateRight(l,4)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

双指针直接找到新的链表的头节点和链接节点。
链接节点的下一个指向老的头，同时老的头需要断链

时间复杂度：O(n)
空间复杂度：O(1)
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if k==0||head == nil {
		return head
	}

	s := head
	f := head
	for i:=0;i<k;i++{
		if f.Next == nil {
			f = head
		} else {
			f = f.Next
		}
	}

	for f.Next != nil {
		s = s.Next
		f = f.Next
	}

	if s.Next == nil {
		return head
	}

	newL := s.Next
	s.Next = nil
	f.Next = head
	return newL
}


func rotateRightOK(head *ListNode, k int) *ListNode {
	if k==0||head == nil {
		return head
	}

	count := 0
	cur := head
	for cur != nil {
		count ++
		cur = cur.Next
	}

	k = k % count

	if k == 0 {
		return head
	}

	s := head
	f := head
	for i:=0;i<k;i++{
		f = f.Next
	}

	for f.Next != nil {
		s = s.Next
		f = f.Next
	}

	if s.Next == nil {
		return head
	}

	newL := s.Next
	s.Next = nil
	f.Next = head
	return newL
}






