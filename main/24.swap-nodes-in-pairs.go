package main

import "fmt"

/**

题目:
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:
给定 1->2->3->4, 你应该返回 2->1->4->3.

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := buildList([]int{1, 2, 3, 4})
	r := swapPairs(l)
	fmt.Println(r, r.Next, r.Next.Next, r.Next.Next.Next)
}

func buildList(nums []int) *ListNode {
	if len(nums) <= 0 {
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

化简问题 x->1->2->3
想要每逢2个交换，就必须
1.将1的下一个2先存起来
2.1指向2的下一个
3.2指向1
4.如果1之前存在节点,就必须将上一个节点的next指向2
5.下一次开始时候，3即为即将的新头，上一个节点变为2

时间复杂度：O(n)
空间复杂度：O(1)
*/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newList *ListNode
	newList = head.Next
	var lastNode *ListNode
	for head != nil && head.Next != nil {
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = head

		if lastNode != nil {
			lastNode.Next = tmp
		}

		lastNode = head
		head = head.Next

	}

	return newList
}

/**
递归的思路
1.找到终止条件：当遍历到最后一个，或者最后一个的下一个为nil的时候（分别应对奇偶的情况）
2.寻找返回值。返回值即为已经完成交换的链表
3.简化交换的逻辑，假设仅剩3个节点的链表进行交换。

1->2->x
需要做的事
将1的下一个暂存起来。
将1指向x，也就是原先2的下一个指向的地方，这里是递归点
将2的下一个指向1
*/
func swapPairsOK(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tmp := head.Next
	head.Next = swapPairsOK(tmp.Next)
	tmp.Next = head
	return tmp
}
