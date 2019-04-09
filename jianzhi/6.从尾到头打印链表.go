package main

import "fmt"

/**
题目：
从尾到头打印链表

分析：
先放入一个栈，然后在进行打印
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
	l := buildList([]int{1, 2, 3, 4, 5})
	printListInReversedList(l)
}

/**

 */
func printListInReversedList(l *ListNode) {
	if l == nil {
		return
	}

	stack := make([]int, 0)
	for l != nil {
		stack = append(stack, l.Val)
		l = l.Next
	}

	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Println(stack[i])
	}
	return
}
