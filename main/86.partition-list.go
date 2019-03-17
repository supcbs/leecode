package main

import (
	"fmt"
)

/**

题目:

给定一个链表和一个特定值 x，对链表进行分隔，
使得所有小于 x 的节点都在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

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
	l:=buildList([]int{1,1})
	r := partition(l,3)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

三指针法：
需要一个指针p表记住小于k的位置，一开始可能是nil
需要一个指针prev记住当前扫描的指针的前一个位置
需要一个指针记住当前扫描的位置

然后将当前的下一个给到prev的下一个
当前的下一个指向p的下一个
p的下一个指向当前的位置

一轮循环后

如果发生了移动则
p移到当前位置的下一个
prev保持不动
cur移到prev的下一个

如果没有移动则
p保持不动
prev移到当前的位置
当前的位置cur后移一个
----
上面的复杂了

直接来两个链表，标记大于和小于的。后面拼接即可


时间复杂度：O(n)
空间复杂度：O(1)
*/
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	newL1 := &ListNode{}
	newL2 := &ListNode{}
	p1 := newL1
	p2 := newL2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			head = head.Next
			p1 = p1.Next
			p1.Next = nil
		} else {
			p2.Next = head
			head = head.Next
			p2 = p2.Next
			p2.Next = nil
		}
	}

	p1.Next = newL2.Next
	return newL1.Next
}
