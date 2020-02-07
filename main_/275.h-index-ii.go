package main

import "fmt"

/**

题目:
给定一位研究者论文被引用次数的数组（被引用次数是非负整数），数组已经按照升序排列。编写一个方法，计算出研究者的 h 指数。
h 指数的定义: “h 代表“高引用次数”（high citations）,
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。
（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）"

(h_index代表有h篇论文的引用次数至少为h次
N 篇论文中，有h篇每篇被引用次数大于等于h，剩下的N - h 篇每篇被引用次数小于等于h。 )


示例:
输入: citations = [0,1,3,5,6]
输出: 3

解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 0, 1, 3, 5, 6 次。
     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。
*/

func main() {
	nums := []int{0, 2, 2, 5, 6}
	r := hIndexII(nums)
	fmt.Println(r)
}

/**
方法1:

单纯二分法，当mid值越小时，此时的H指数将越大

时间复杂度：O(logn)
空间复杂度：O(1)
*/
func hIndexII(citations []int) int {
	citationsLen := len(citations)
	var res int

	left := 0
	right := citationsLen - 1
	for left <= right {
		mid := (right + left) >> 1
		fmt.Println(citations[mid], citationsLen, mid)
		if citations[mid] >= citationsLen - mid {
			res = citationsLen - mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
