package main

import "fmt"

/**

题目:

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := &ListNode{
		Val: 2,
	}
	deleteNode(n)

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

/**
方法1:

关键在与遍历链表的时候记住上一个。
如果当前值等于要删除的值，则直接将上一个的next设置为当前的next即可

时间复杂度：O(n)
空间复杂度：O(1)
*/
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
