package main


/**
双指针，快的先走n步，然后慢的和快的一起走，直到重点。此时注意，下一个节点即为要删除的点
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