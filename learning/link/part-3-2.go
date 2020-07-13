package main

/**
移除指定元素
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	newHead := &ListNode{
		Next:head,
	}
	prev := newHead
	for head != nil {
		if head.Val == val {
			prev.Next = head.Next
		} else {
			prev = head
		}
		head = head.Next
	}

	return newHead.Next
}
