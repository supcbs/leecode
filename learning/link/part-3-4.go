package main

/**
关键点在于：快慢指针判断中间位置，然后顺道翻转慢指针之前的链表。

有两个不同的地方：
1.这里的翻转不是要head，也就是只需要记住上一个即可
2.判断奇偶可以根据快指针的来判断，如果最好快指针有值则为奇数
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	var pre,cur *ListNode

	for fast != nil && fast.Next != nil {
		cur = slow
		slow = slow.Next
		fast = fast.Next.Next

		cur.Next = pre
		pre = cur
	}

	if fast != nil {
		slow = slow.Next
	}

	for cur != nil {
		if slow.Val != cur.Val {
			return false
		}
		cur = cur.Next
		slow = slow.Next
	}

	return true
}
