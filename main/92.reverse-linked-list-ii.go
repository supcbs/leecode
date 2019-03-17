package main

import "fmt"

/**

题目:

反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

*/

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l:=buildList([]int{1,2,3,4,5})
	r := reverseBetween(l,1,2)
	fmt.Println(r,r.Next,r.Next.Next,r.Next.Next)
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	root := &ListNode{
		Val:nums[0],
	}
	tmp := root
	for i:=1;i<len(nums);i++{
		tmp.Next = &ListNode{
			Val:nums[i],
		}
		tmp = tmp.Next
	}

	return root
}

/**
方法1:

首先将指针指到m的位置，然后开始链表翻转。
最后在拼接三段

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var prev,tail,con *ListNode
	cur := head
	for m > 1 {
		prev = cur
		cur = cur.Next
		m--
		n--
	}

	var newL *ListNode
	tail = cur
	con = prev
	for n > 0 {
		next := cur.Next
		cur.Next = newL
		newL = cur
		cur = next
		n--
	}

	if con != nil {
		con.Next = newL
	} else {
		head = newL
	}

	tail.Next = cur //  cur已经指向了下一个
	return head
}


