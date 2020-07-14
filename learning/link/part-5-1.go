package main


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	newHead := &ListNode{}
	p := newHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}

		p = p.Next
	}

	for l1 != nil {
		p.Next = l1
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		p.Next = l2
		p = p.Next
		l2 = l2.Next
	}

	return newHead.Next
}