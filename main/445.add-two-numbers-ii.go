package main

import "fmt"

/**
题目:
给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。
你可以假设除了数字 0 之外，这两个数字都不会以零开头。

进阶:
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:
输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7

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
	l1 := buildList([]int{7, 2, 4, 3})
	l2 := buildList([]int{5, 6, 4})
	r := addTwoNumbersiiT(l1, l2)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:
现将两个链表变为栈，然后在一个一个弹出相加

边界注意：
1.长短不一
2.可能最后存在进位

时间复杂度：O(n)
空间复杂度：O(n)
*/
func addTwoNumbersiiT(l1 *ListNode, l2 *ListNode) *ListNode {
	m1, m2 := make([]int, 0), make([]int, 0)
	p1, p2 := l1, l2

	for p1 != nil {
		m1 = append(m1, p1.Val)
		p1 = p1.Next
	}

	for p2 != nil {
		m2 = append(m2, p2.Val)
		p2 = p2.Next
	}

	var head *ListNode
	n1, n2 := len(m1)-1, len(m2)-1
	carry := int(0)
	for n1 >= 0 || n2 >= 0 || carry > 0 {
		sum := carry
		if n1 >= 0 {
			sum += m1[n1]
		}

		if n2 >= 0 {
			sum += m2[n2]
		}

		if sum > 9 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}

		node := &ListNode{
			Val:  sum,
			Next: head,
		}

		head = node

		n1--
		n2--
	}

	if carry > 0 {
		node := &ListNode{
			Val:  carry,
			Next: head,
		}

		head = node
	}

	return head
}

/**
下面是跑得最快的，但是这么些可能会被要求不能反转数组
 */
func addTwoNumbersiiFast(l1 *ListNode, l2 *ListNode) *ListNode {
	reverl1 := reverse(l1)
	reverl2 := reverse(l2)
	var temp *ListNode
	var newHead *ListNode
	up := 0
	for {
		if reverl1 == nil && reverl2 == nil && up == 0{
			break
		}
		newNode,newUp := add(reverl1,reverl2,up)
		up = newUp
		if temp != nil {
			temp.Next = newNode
		}else{
			newHead = newNode
		}
		temp = newNode
		if reverl1 != nil{
			reverl1 = reverl1.Next
		}
		if reverl2 != nil{
			reverl2 = reverl2.Next
		}
	}
	return reverse(newHead)
}

func add(l1 *ListNode,l2 *ListNode,up int) (*ListNode,int){
	var res int
	var newUp int
	if l1 == nil && l2 == nil{
		res = up
		newUp = 0
	}else if l1 == nil {
		res = (l2.Val + up) % 10
		newUp = (l2.Val + up) / 10
	}else if l2 == nil {
		res = (l1.Val + up) % 10
		newUp = (l1.Val + up) / 10
	}else{
		res = (l1.Val + l2.Val + up) % 10
		newUp = (l1.Val + l2.Val + up) / 10
	}
	return &ListNode{res,nil},newUp
}

func reverse(head *ListNode) *ListNode{
	var front *ListNode
	var next *ListNode
	for {
		if head == nil{
			break
		}
		next = head.Next
		head.Next = front
		front = head
		head = next
	}
	return front
}