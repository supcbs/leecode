package main

import (
	"fmt"
)

/**

题目:
给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），
可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:
输入: [1,3,4,2,2]
输出: 2

示例 2:
输入: [3,1,3,4,2]
输出: 3

说明：
不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。

*/

func main() {
	nums := []int{3,1,3,4,2}
	r := findDuplicate2(nums)
	fmt.Println(r)
}

/**
方法1:
说明已经限制掉了很多方法，比如存hash等
由于总共只会有n+1个，并且数组的数范围在1-n直接，所以必定有一个数重复。
可以采用二分法进行排除。

每次取出left和right的数除二取中位数mid，计算小于等于中位数的个数num，
如果大于mid，说明重复的数在左边，将right置为mid，反之右边
继续计算。

时间复杂度：O(nlogn)
空间复杂度：O(1)
*/
func findDuplicate(nums []int) int {
	numsLen := len(nums)

	if numsLen == 2 {
		return nums[0]
	}

	left := 0
	right := numsLen - 1
	for left < right {
		mid := (left + right) / 2

		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}

		// 这里的边界条件要想清楚：上头是<=mid，下头当>mid的时候，必须left置为mid+1
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right
}

/**
方法1:
由于题目的特殊性，反之肯定跳不出n的长度，可以通过快慢指针的方式进行处理
slow = num[i]
fast = num[num[i]]

先按照快慢的方式，得到相遇的点。将快的置为初始位置，然后再次开始跑，相遇的点就是重复的数

数量小可以但是数量大的抛锚

时间复杂度：O(n)？？
空间复杂度：O(1)
*/
func findDuplicate2(nums []int) int {
	numsLen := len(nums)

	if numsLen < 2 {
		return 0
	}
	if numsLen == 2 {
		return nums[0]
	}

	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return fast
}
