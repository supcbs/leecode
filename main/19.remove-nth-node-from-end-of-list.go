package main

import "fmt"

/**

题目:
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l:=buildList([]int{1,2,3,4,5})
	//r := removeNthFromEnd2(l,2)
	l := buildList([]int{1})
	r := removeNthFromEnd(l, 1)
	fmt.Println(r) //,r.Next,r.Next.Next,r.Next.Next.Next)
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

/**
方法1:

j先走n-1步，然后ij指针一起走，直到j到尾巴，i的位置就是需要删除的

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 0 {
		return nil
	}

	fast := &ListNode{
		Next: head,
	}
	slow := &ListNode{
		Next: head,
	}

	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 有可能只有一个节点，且删除第一个，那么此时slow实际上一步没走
	if slow.Next == head {
		head = head.Next
		return head
	}

	slow.Next = slow.Next.Next
	return head
}
