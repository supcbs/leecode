package main

import "fmt"

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

func main(){
	head := buildList([]int{1, 2})
	r := reverseList(head)
	fmt.Println(r,r.Next)
}

// 反转链表
func reverseList(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	var newL *ListNode
	for node != nil {
		tmp := node.Next
		node.Next = newL
		newL = node
		node = tmp
	}

	return newL
}


// 排序
// 求交点
// 倒K个节点
// 每隔K隔反转
