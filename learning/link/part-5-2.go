package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	count := 0
	newHead := &ListNode{}
	p := newHead
	for l1 != nil || l2 != nil || count > 0 {
		if l1 != nil {
			count += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			count += l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{
			Val: count % 10,
		}
		p = p.Next
		count /= 10
	}

	return newHead.Next
}
