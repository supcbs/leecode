package main

import "fmt"

/**

题目:

*/

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l:=buildList([]int{1,1,3,4,4})
	r := deleteDuplicates(l)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
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

时间复杂度：O(n)
空间复杂度：O(1)
*/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}

	var h ListNode
	h.Next = head
	for head.Next != nil {
		fmt.Println("headh",head)
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}
	}

	return h.Next
}

/**
if head == nil || head.Next == nil{
		return head
	}

	i := head
	j := head.Next
	for j != nil {
		fmt.Println("2",i,j)
		if i.Val == j.Val {
			i.Next = j
			i = j
		}
		j = j.Next
	}

	//if j.Val == i.Val {
	//	i.Next = j.Next
	//}
 */
