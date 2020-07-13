package main

/**
两种方法：
1. 将其中一个链表的尾部指向头部，然后从head遍历另一个链表求是否有环。变成part-2-2的那个题
2. 跑链表a和链表b，当跑完的时候将指针换到另一个链表，这样如果有交点就一定会相遇，且步数不会超过两个链表加在一起的长度。（这里写一下这个）
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	count := 0
	for h:=headA;h!= nil;h=h.Next{
		count++
	}
	for h:=headB;h!= nil;h=h.Next{
		count++
	}

	p1 := headA
	p2 := headB
	for p1 != p2 && count > 0 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}

		if p2 == nil {
			p2 = headA
		} else {
			p2 = p1.Next
		}
		count--
	}

	if p1 != p2 {
		return nil
	}

	return p1
}