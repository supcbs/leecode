package main

import "fmt"

/**

题目:

给定一个链表，判断链表中是否有环。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。


示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：true
解释：链表中有一个环，其尾部连接到第二个节点。


示例 2：
输入：head = [1,2], pos = 0
输出：true
解释：链表中有一个环，其尾部连接到第一个节点。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := buildList()
	r := hasCycle(nil)
	fmt.Println(r)
}

/**
方法1:

判断环：
一开始都指向头，两个指针分别向前，i以1的速度，j以2的速度，当i所指=j所指则认为没有环

关键点在于终止条件，需要是快的不为nil，并且快的下一个也不为nil（保证了两次下一个都不是nil）

时间复杂度：O(n)
空间复杂度：O(1)
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
