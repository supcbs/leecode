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
	l := buildList([]int{1, 2, 3,5,3, 2,1})
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