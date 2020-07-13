package main


/**

head
1->2->3->nil

tmp = 2

head->newL
newL = head // 相当于放在他前面了

head = tmp
 */

/**
leetcode 206
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newList *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = newList
		newList = head
		head = tmp
	}

	return newList
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

