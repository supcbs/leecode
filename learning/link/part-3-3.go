package main

/**
关键点在于：奇偶需要单独存，同时记住偶数的头，后面放到技术尾巴
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	j := head
	o := head.Next
	tmp := head.Next

	for j.Next != nil && j.Next.Next != nil {
		j.Next = o.Next
		j = j.Next
		o.Next = j.Next
		o = o.Next
	}

	j.Next = tmp
	return head
}
