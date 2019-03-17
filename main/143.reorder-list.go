package main

import "fmt"

/**

题目:

给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.

示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	tmp := root
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}

	return root
}

func main() {
	ll := buildList([]int{1, 4, 2, 3})
	l := ll
	reorderList(ll)
	for l != nil {
		fmt.Println(l)
		l = l.Next
	}
}

/**
方法1:

方法就是
1.快慢指针区分出前后两部分
2.将慢指针的下一个存tmp后，进行断链（next=nil），将tmp之后的进行翻转。
3.遍历翻转的链表，进行夹缝插入

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	//1.快慢指针区分出前后两部分
	s := head
	f := head
	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	// 将慢指针的下一个存tmp后，进行断链（next=nil），将tmp之后的进行翻转。
	needReverse := s.Next
	s.Next = nil

	var newL *ListNode
	p := newL
	for needReverse != nil {
		tmp := needReverse.Next
		needReverse.Next = p
		p = needReverse
		needReverse = tmp
	}

	fmt.Println(needReverse,s,p)

	//3.遍历翻转的链表，进行夹缝插入
	for p != nil {
		curT := head.Next
		head.Next = p
		needT := p.Next
		p.Next = curT

		head = curT
		p = needT
	}

}
