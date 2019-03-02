package main

import (
	"fmt"
	"container/list"
)

/**

题目:
给定一个数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口 k 内的数字。滑动窗口每次只向右移动一位。

返回滑动窗口最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]

解释:

  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
注意：

你可以假设 k 总是有效的，1 ≤ k ≤ 输入数组的大小，且输入数组不为空。

进阶：

你能在线性时间复杂度内解决此题吗？

*/

func main() {
	nums := []int{1, 3,-1,-3,5,3,6,7}
	k:=3
	//nums = []int{1}
	//k=1
	r := maxSlidingWindowOK(nums,k)
	fmt.Println(r)
}

/**
方法1:

两个指针搞定。

时间复杂度：O(n + ) ？
空间复杂度：O(1)
*/
func maxSlidingWindow(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen <= 0 || numsLen < k{
		return []int{}
	}

	list.New()

	ret := []int{}
	intMin := ^int(^uint(0) >> 1)

	p1 := 0
	p2 := k-1
	for ;p2 < numsLen;p2++{
		max := intMin
		for i := p1;i<=p2;i++{
			if nums[i] > max {
				max = nums[i]
			}
		}

		ret = append(ret, max)
		p1++
	}


	return ret
}


/**
方法2:

大神使用链表。
当i>=k的时候，首先保证链表中的i是属于窗口的，即：小于i-k的要被干掉
然后删除当前i对应的val小的链表节点
此时再加入该节点的i,存在两种情况：
1.都比它小，则他就是最大的
2.存在部分比它小，则它排在窗口后面（比如说老二的位置）

当i>k的时候则认为认为排在链表最前面的就是最大的，加到返回值中

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxSlidingWindowOK(nums []int, k int) []int {
	numsLen := len(nums)
	if numsLen <= 0 || numsLen < k{
		return []int{}
	}

	window := list.New()
	var ret []int

	for i,v :=  range nums {
		// 删除有两个条件：
		// 1.i的位置已经到达了k，说明有的删了
		// 2.如果链表最前头的索引index，比窗口的前头小，则需要删掉
		if i>=k && window.Front().Value.(int) <= i-k {
			window.Remove(window.Front())
		}

		// 与当前的val比，仅留下一个值最大的索引
		for window.Len() > 0 && nums[window.Back().Value.(int)] <= v {
			window.Remove(window.Back())
		}

		// 这里讲究了，如果当前最大，则i肯定是最大的，因为前面全删了
		// 如果不比前几个大，那么前一步只会把比他小的干掉，最大的依旧在最前面
		window.PushBack(i)
		// 从k-1开始，可以开始增就返回值了
		if i >= k - 1 {
			fmt.Println("add ret:",window.Front().Value.(int))
			ret = append(ret, nums[window.Front().Value.(int)])
		}
	}

	return ret
}


