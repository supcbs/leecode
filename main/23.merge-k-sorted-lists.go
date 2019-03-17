package main

import (
	"fmt"
)

/**

题目:

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
示例:
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

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
	l1:=buildList([]int{1,4,5})
	l2:=buildList([]int{1,3,4})
	l3:=buildList([]int{2,6})
	lArr := []*ListNode{l1,l2,l3}
	r := mergeKListsOK(lArr)
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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	var tmp *ListNode
	tmp = mergeSortLists(lists[0],lists[1])
	for i:=2;i<len(lists);i++{
		tmp = mergeSortLists(tmp,lists[i])
	}

	return tmp
}

/**
这道题也是可以用归并排序的
 */
func mergeKListsOK(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	if len(lists) == 2 {
		return mergeSortLists(lists[0],lists[1])
	}

	listLen := len(lists)
	left := mergeKListsOK(lists[:listLen/2])
	right := mergeKListsOK(lists[listLen/2:])

	return mergeSortLists(left,right)
}

func mergeSortLists(l1,l2 *ListNode) *ListNode{
	if l1 == nil && l2 == nil {
		return nil
	}

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

	if l1 != nil {
		p.Next = l1
	} else {
		p.Next = l2
	}

	return newL.Next
}
