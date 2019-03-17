package main

import (
	"fmt"
)

/**

题目:
请判断一个链表是否为回文链表。

示例 1:
输入: 1->2
输出: false

示例 2:
输入: 1->2->2->1
输出: true

进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val: nums[0],
	}
	tmp := root
	for i := 1; i < len(nums); i++ {
		tmp.Next = &ListNode{
			Val: nums[i],
		}
		tmp = tmp.Next
	}

	return root
}

func main() {
	l := buildList([]int{1, 2, 3,3, 2,1})
	r := isPalindrome(l)
	fmt.Println(r)
}

/**
方法1:

这道题贼恶心

1,2,3,4,3,2,1

1.先利用快慢指针让快指针到尾巴。
需要注意的是快指针如果是奇数则nil的下一步不是没有的，同判断出奇偶
2.然后开始翻转head开始到slow的位置[3-2-1]（此时slow是过了一半的第一个位置）
翻转回来的链表4-3-2-1
3.如果此时是奇数，则翻转的链表需要往后进一位再进行比较
4.比较q和slow即可

时间复杂度：O(n)
空间复杂度：O(1)
*/

func isPalindrome(head *ListNode) bool {
	slow,fast := head,head
	isOdd := false

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		} else {
			isOdd = true
		}
		fmt.Println(slow,fast)
	}

	q:=reverse(head, slow)

	if isOdd && q!=nil {
		q = q.Next
	}
	fmt.Println("12-",q,head,slow)
	//for q!=nil{
	//	q=q.Next
	//	fmt.Println(q)
	//}
	for slow != nil {
		if slow.Val != q.Val {
			return false
		}
		slow = slow.Next
		q = q.Next
	}

	return true
}
func reverse (head *ListNode, tail *ListNode) *ListNode {
	var p *ListNode
	q := head
	for q != nil && q!=tail {
		tmp := q.Next
		q.Next = p
		p = q
		q = tmp
	}
	return p
}
