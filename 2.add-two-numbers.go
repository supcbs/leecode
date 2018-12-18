package main

import (
	"fmt"
)

/**

题目:

给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//l1 := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 4,
	//		Next: &ListNode{
	//			Val:  3,
	//			Next: nil,
	//		},
	//	},
	//}
	//
	//l2 := &ListNode{
	//	Val: 5,
	//	Next: &ListNode{
	//		Val: 6,
	//		Next: &ListNode{
	//			Val:  4,
	//			Next: nil,
	//		},
	//	},
	//}

	//l1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 8,
	//		Next: nil,
	//	},
	//}
	//
	//l2 := &ListNode{
	//	Val: 0,
	//	Next: nil,
	//}

	l1 := &ListNode{
		Val: 5,
		Next: nil,
	}

	l2 := &ListNode{
		Val: 5,
		Next: nil,
	}

	l3 := addTwoNumbers(l2, l1)
	fmt.Println(l3.Val, *l3.Next, (*l3.Next).Val)
	//fmt.Println(l3.Val, *l3.Next, (*l3.Next).Val, *(*l3.Next).Next, (*(*l3.Next).Next).Val, (*(*l3.Next).Next).Next)
}

// 1. 不用递归
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	returnl := &ListNode{}
	var x, y, sum int
	p, q, carry := l1, l2, 0
	temp := returnl
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}

		if q != nil {
			y = q.Val
		} else {
			y = 0
		}

		sum = x + y + carry
		carry = sum / 10
		temp.Val = sum % 10
		if (p == nil || p.Next == nil) && (q == nil || q.Next == nil) {
			break
		}
		temp.Next = &ListNode{}
		temp = temp.Next

		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}

	if carry > 0 {
		temp.Next = &ListNode{1,nil}
	}

	return returnl
}


// 2.使用递归
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = &ListNode{}
	}

	if l2 == nil {
		l2 = &ListNode{}
	}

	result := &ListNode{}
	result.Val = l1.Val + l2.Val

	if result.Val > 9 {
		result.Val -= 10
		if l1.Next == nil {
			l1.Next = &ListNode{}
		}
		l1.Next.Val += 1
	}

	if l1.Next == nil && l2.Next == nil {
		return result
	}

	result.Next = addTwoNumbers(l1.Next, l2.Next)
	return result
}

