package main

import "fmt"

/**

题目:
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:
输入: 4->2->1->3
输出: 1->2->3->4

示例 2:
输入: -1->5->3->4->0
输出: -1->0->3->4->5

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
	l:=buildList([]int{1,3,2})
	r := sortList(l)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:


 * 参考：Sort List——经典（链表中的归并排序） https://www.cnblogs.com/qiaozhoulin/p/4585401.html
 *
 * 归并排序法：在动手之前一直觉得空间复杂度为常量不太可能，因为原来使用归并时，都是 O(N)的，
 * 需要复制出相等的空间来进行赋值归并。对于链表，实际上是可以实现常数空间占用的（链表的归并
 * 排序不需要额外的空间）。利用归并的思想，递归地将当前链表分为两段，然后merge，分两段的方
 * 法是使用 fast-slow 法，用两个指针，一个每次走两步，一个走一步，知道快的走到了末尾，然后
 * 慢的所在位置就是中间位置，这样就分成了两段。merge时，把两段头部节点值比较，用一个 p 指向
 * 较小的，且记录第一个节点，然后 两段的头一步一步向后走，p也一直向后走，总是指向较小节点，
 * 直至其中一个头为NULL，处理剩下的元素。最后返回记录的头即可。
 *
 * 主要考察3个知识点，
 * 知识点1：归并排序的整体思想
 * 知识点2：找到一个链表的中间节点的方法
 * 知识点3：合并两个已排好序的链表为一个新的有序链表


时间复杂度：O(n)
空间复杂度：O(1)
*/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 找到中心节点
	fast := head
	slow := head

	// 多一个fast.Next.Next，是想让分半的时候，当奇数，慢指针多泡一个
	for fast.Next !=nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	left := head
	right := slow.Next
	slow.Next = nil

	// 进行归并排序
	leftPart := sortList(left)
	rightPart := sortList(right)

	return mergeSortList(leftPart,rightPart)
}

func mergeSortList(l1,l2 *ListNode) *ListNode {
	newL := &ListNode{}
	p := newL
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
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

	return newL.Next
}
