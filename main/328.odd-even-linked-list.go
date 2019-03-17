package main

import "fmt"

/**

题目:
给定一个单链表，把所有的奇数节点和偶数节点分别排在一起。
请注意，这里的奇数节点和偶数节点指的是节点编号的奇偶性，而不是节点的值的奇偶性。
请尝试使用原地算法完成。
你的算法的空间复杂度应为 O(1)，时间复杂度应为 O(nodes)，nodes 为节点总数。

示例 1:
输入: 1->2->3->4->5->NULL
输出: 1->3->5->2->4->NULL

示例 2:
输入: 2->1->3->5->6->4->7->NULL
输出: 2->3->6->7->1->5->4->NULL
说明:

应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。

*/

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l := buildList([]int{1,2,3,4,5})
	r := oddEvenList(l)
	fmt.Println(r,r.Next,r.Next.Next,r.Next.Next.Next)
}

func buildList(nums []int) *ListNode {
	if len(nums) <= 0 {
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

终止条件：当前奇数的下一个或者当前奇数为nil
核心思路是：
1.奇数第一个和偶数第一个先选出,暂存偶数第一个后面用于拼接在奇数的后面
2.奇数的下一个 = 偶数的下一个
3.偶数的下一个 = 奇数的下两个（此时奇数的下一个已经变为原先偶数的下一个了，所以这里的下两个）


时间复杂度：O(n)
空间复杂度：O(1)
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	j := head
	o := head.Next
	tmp := head.Next
	for j.Next != nil && j.Next.Next != nil {
		j.Next = o.Next
		j = j.Next
		o.Next = j.Next
		o = o.Next
	}

	j.Next = tmp
	return head
}
