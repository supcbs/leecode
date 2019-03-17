package main

import (
	"fmt"
)

/**

题目:
编写一个程序，找到两个单链表相交的起始节点。

如下面的两个链表：

在节点 c1 开始相交。

示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。


示例 2：
输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Reference of the node with value = 2
输入解释：相交节点的值为 2 （注意，如果两个列表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。


示例 3：
输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
输入解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
解释：这两个链表不相交，因此返回 null。


注意：
如果两个链表没有交点，返回 null.
在返回结果后，两个链表仍须保持原有的结构。
可假定整个链表结构中没有循环。
程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(nums []int, l *ListNode) *ListNode {
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

	tmp.Next = l
	return root
}

func main() {
	l := buildList([]int{1,2,3}, nil)
	l1 := buildList([]int{4,6}, l)
	l2 := buildList([]int{5}, l)
	r := getIntersectionNodeNice(l1, l2)
	for r != nil {
		fmt.Println(r)
		r = r.Next
	}
}

/**
方法1:

两个指针一起跑。当跑完当前链表的时候，指向另一个链表的头部，相等的时候即为交点.
这个其实相当于拿两个链表相互对比了

时间复杂度：O(n)
空间复杂度：O(1)
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p1 := headA
	p2 := headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}

	if p1 != p2 {
		return nil
	}
	return p1
}

/**
A：5-1-2-3
B：4-6-1-2-3
看到一种很妙的思路，先将a变成一个环。

然后对于B来说就是： 4-6-1-2-3-5
					  |     |
                      -------
这样就变成了找环的交点问题
*/

func getIntersectionNodeNice(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	last := headB
	for last.Next != nil {
		last = last.Next
	}

	last.Next = headB

	fast := headA
	slow := headA
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		fmt.Println("--",slow,fast)
		if slow == fast {
			fmt.Println(slow, fast)
			// 解环
			slow = headA
			fmt.Println(slow, fast)
			for slow!=fast {
				slow = slow.Next
				fast = fast.Next
			}
			last.Next = nil
			return fast
		}
	}

	return nil
}
