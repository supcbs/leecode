package main

import "fmt"

/**

题目:
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。

示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。

示例 2：
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。

示例 3：
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。


进阶：
你是否可以不用额外空间解决此题？

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
	l:=buildList([]int{1,2,3,4,5})
	r := detectCycle(l)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

1。快慢指针判断是否有环
2。存在则慢指针从头开始，然后和快指针一起一步一步向后，相遇的点就是交点

时间复杂度：O(n)
空间复杂度：O(1)
*/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	s := head
	f := head
	for f!=nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			s = head
			for s != f {
				s = s.Next
				f = f.Next
			}

			return s
		}
	}

	return nil
}