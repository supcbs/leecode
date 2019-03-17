package main

import "fmt"

/**

题目:

*/

type ListNode struct {
	Val int
	Next *ListNode
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

func main() {
	l:=buildList([]int{1,2,3,4,5})
	r := temp1(l)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func temp1(head *ListNode) *ListNode {
	return nil
}
