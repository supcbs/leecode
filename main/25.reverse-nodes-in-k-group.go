package main

import (
	"fmt"
)

/**

题目:

给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。
k 是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :
给定这个链表：1->2->3->4->5
当 k = 2 时，应当返回: 2->1->4->3->5
当 k = 3 时，应当返回: 3->2->1->4->5

说明 :
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

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
	l := buildList([]int{1, 2, 3, 4, 5, 6, 7, 8})
	r := reverseKGroup(l, 3)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

递归三部曲：
1.边界条件：当当前的头不满足k个的时候进行返回
2.寻找返回值，返回值就是已经翻转好的链表。
这么想：已经翻转的链表 接上即将需要翻转的链表 = 最终链表。这里的即将需要翻转的链表就是递归的点
3.问题简化为翻转链表的头三个节点即可

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	tmpNode := head

	// 判断是否满足三个
	for i := 0; i < k; i++ {
		if tmpNode == nil {
			return head
		}
		tmpNode = tmpNode.Next
	}

	// 未来的尾节点
	tmpNode = head
	for i := 0; i < k-1; i++ {
		/**
		这里移动的思路：将tmp的下一个移到head位置
		1234
		2134
		3214
		 */
		N := tmpNode.Next
		tmpNode.Next = N.Next
		N.Next = head
		head = N
		// 将当前节点的下一个节点移到前头
		//N2 := tmpNode.Next
		////tmpNode.Next = N2.Next
		////head = N2
		////N2.Next = head
		//tmpNode.Next, head, N2.Next = N2.Next, N2, head
		//fmt.Println("--",tmpNode, N2,head)

		//tNext := tmpNode.Next
		//tNextNext := tNext.Next
		//head.Next = tNextNext
		//tNext.Next = head
		//head = tNext
	}
	tmpNode.Next = reverseKGroup(tmpNode.Next, k)
	fmt.Println("tmp", tmpNode)
	return head
}


func reverseKGroupOK(head *ListNode, k int) *ListNode {
	tmpNode := head
	for i := 0; i < k; i++ {
		if tmpNode == nil {
			return head
		}
		tmpNode = tmpNode.Next
	}

	tmpNode = head
	for i := 0; i < k-1; i++ {
		// move N2 to head
		N2 := tmpNode.Next
		tmpNode.Next, head, N2.Next = N2.Next, N2, head
	}
	tmpNode.Next = reverseKGroup(tmpNode.Next, k)

	return head
}