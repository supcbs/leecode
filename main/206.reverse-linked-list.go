package main

import (
	"fmt"
)

/**

题目:
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := buildList([]int{1, 2, 3, 4, 5})
	r := reverseListOK(head)
	fmt.Println(r,r.Next)
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

翻转链表的核心在于:
1.首先要把下一个节点给存下来，以免断链
2.将当前节点的下一个节点next设置为上一个节点prev
3.将当前节点设置为上一个节点
4.将下一个的当前节点设置为当前节点的下一个节点，

停止时机：当下一个节点为nil时候，那么将这个节点取出就是翻转过后新链表的头

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseList(head *ListNode) *ListNode {
	var newHead *ListNode
	var prevNode *ListNode
	curNode := head
	for curNode != nil {
		nodeNext := curNode.Next
		if nodeNext == nil {
			newHead = curNode
		}
		curNode.Next = prevNode

		prevNode = curNode
		curNode = nodeNext
	}

	return newHead
}

/**
头节点插入法

 */
func reverseListOK(head *ListNode) *ListNode {
	var newL *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = newL
		newL = head
		head = tmp
	}
	fmt.Println(newL,newL.Next)
	return newL
}
/**
https://blog.csdn.net/fx677588/article/details/72357389
 */
func reverseListOK2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseListOK2(head.Next)
	// 翻转链表的指向
	head.Next.Next = head

	// 断开下一个的链接
	head.Next = nil
	return newHead
}
