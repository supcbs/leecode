package main

import (
	"fmt"
	"sort"
)

/**

题目:
给定一位研究者论文被引用次数的数组（被引用次数是非负整数）。编写一个方法，计算出研究者的 h 指数。
h 指数的定义: “h 代表“高引用次数”（high citations），
一名科研人员的 h 指数是指他（她）的 （N 篇论文中）至多有 h 篇论文分别被引用了至少 h 次。
（其余的 N - h 篇论文每篇被引用次数不多于 h 次。）”

(h_index代表有h篇论文的引用次数至少为h次
N 篇论文中，有h篇每篇被引用次数大于等于h，剩下的N - h 篇每篇被引用次数小于等于h。 有多个符合条件的h的话取最大的那个。)

示例:
输入: citations = [3,0,6,1,5]
输出: 3

解释: 给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
     由于研究者有 3 篇论文每篇至少被引用了 3 次，其余两篇论文每篇被引用不多于 3 次，所以她的 h 指数是 3。
说明: 如果 h 有多种可能的值，h 指数是其中最大的那个。
*/

func main() {
	nums := []int{3,0,6,1,5,4}
	//r := hIndexOK(nums)
	r := hIndex2(nums)
	fmt.Println(r)
}

/**
方法1:
必须满足n篇文章至少被引用n次，则n的取值至多不会大于初始的论文篇数（即数组长度）
通过暂存key为引用次数（大于论文篇数的记为论文篇数），value为每次+1的map
再进行倒着（因为h多个时候取最大的）循环遍历map，判断k==v则ok，
若k!=v,则将v累加到k-1上，（因为大于k了肯定大于k-1,所以累加没问题）

时间复杂度：O(n)
空间复杂度：O(n)
*/
func hIndexOK(citations []int) int {
	cLen := len(citations)
	tempMap := make([]int, cLen+1)
	for _, v := range citations {
		if v > cLen {
			tempMap[cLen]++
		} else {
			tempMap[v]++
		}
	}

	for i := cLen; i > 0; i-- {
		if tempMap[i] >= i {
			return i
		}
		tempMap[i-1] += tempMap[i]
	}
	return 0
}

/**
方法1:
还有一种是先排序，后从大到小遍历，每次遍历index都+1，这样就会很快遍历到n=index的数
（这里要注意的是既然是从大到小比如6遍历完index=1，5遍历完index++后=2，
	既然从大到小，所以完全可以保证遍历过的的大于当前的index）

时间复杂度：O(nlogn)?
空间复杂度：O(1)
*/
func hIndex2(citations []int) int {
	sort.Ints(citations)
	cLen := len(citations)

	var index int
	for i := cLen - 1; i >= 0; i-- {
		if citations[i] > index {
			index++
		}
	}

	return index
}
