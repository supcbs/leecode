package main

import (
	"fmt"
)

/**

题目:

删除链表中等于给定值 val 的所有节点。
示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l := buildList([]int{1, 2, 3, 4, 4})
	r := removeElements(l, 4)
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

由于可能第一个就被过滤掉所以创建一个新头节点
头节点的下一个指向头节点。
prev指向新头
当值相等的时候，prev位指向当前的下一个
不等的时候，移动prev的位置置当前位置

此时head每一步都进一位，这样全程下来就可以保持head都是跑在prev前面

时间复杂度：O(n)
空间复杂度：O(1)
*/
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	// 新建的头节点，因为有可能后面一个没有所有
	var newh ListNode
	newh.Next = head
	prev := &newh
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}

		head = head.Next
	}

	return newh.Next
}
