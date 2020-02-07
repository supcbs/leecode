package main

import "fmt"

/**

题目:
给出两个 非空 的链表用来表示两个非负的整数。
其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

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
	l1:=buildList([]int{2,4,3})
	l2:=buildList([]int{5,6,4})
	r := addTwoNumbersOK(l1,l2)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

使用递归进行处理。
关键点在于超过10的时候，需要进一位，下一位不存在的话需要进行初始化。

还有一个点就是有可能l1和l2不一样长，所以在这种情况下每次进行加减的时候需要进行初始化

时间复杂度：O(n)
空间复杂度：O(n)
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		l1 = &ListNode{}
	}

	if l2 == nil {
		l2 = &ListNode{}
	}

	result := &ListNode{}
	result.Val = l1.Val + l2.Val

	if result.Val > 9 {
		result.Val %= 10
		if l1.Next == nil{
			l1.Next = &ListNode{}
		}
		l1.Next.Val +=1
	}

	if l1.Next == nil && l2.Next == nil {
		return result
	}

	result.Next = addTwoNumbers(l1.Next,l2.Next)

	return result
}

/**
非递归的方式
 */
func addTwoNumbersOK(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	h := head
	count := 0
	for l1!=nil||l2!=nil||count!=0 {
		if l1!=nil {
			count += l1.Val
			l1 = l1.Next
		}
		if l2 !=nil {
			count += l2.Val
			l2 = l2.Next
		}

		node := &ListNode {
			Val: count %10,
		}
		head.Next = node
		head = head.Next
		count /=10
	}

	return h.Next
}
