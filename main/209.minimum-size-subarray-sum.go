package main

import (
	"fmt"
	"math"
)

/**

题目:
给定一个含有 n 个正整数的数组和一个正整数 s ，
找出该数组中满足其和 ≥ s 的长度最小的连续子数组。
如果不存在符合条件的连续子数组，返回 0。

示例:
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。

进阶:
如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。

*/

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 7
	r := minSubArrayLen(s, nums)
	fmt.Println(r)
}

/**
方法1:

与76题类型

本题采用滑动窗口 + 双指针的方式解题：
首先定义两个指针start，end都从0位置开始。
end不断向前，直至总和>=目标值，那么算出子数组数量，更新最短子数组数量，
这里有个关键一步：此时在>=s的窗口中做start指针向前跑，直至不满足总和>=s.

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minSubArrayLen(s int, nums []int) int {

	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}

	start := 0
	sum := 0
	minLen := numsLen + 1

	for end := 0; end < numsLen; end++ {
		sum += nums[end]

		for sum >= s {
			tempMinLen := end - start + 1
			if tempMinLen < minLen {
				minLen = tempMinLen
			}
			sum -= nums[start]
			start++
		}
	}

	return minLen %(numsLen + 1)
}

func minSubArrayLen(target int, nums []int) int {
	// 边界条件判断
	if len(nums) <= 0 {
		return 0
	}

	// 初始值设置
	i, j := 0, 0 // 滑动指针
	minLen := math.MaxInt32 // 子数组最小长度
	sum := 0 // 子数组的和

	// 滑动窗口实现
	for j < len(nums) {
		sum += nums[j]
		if sum >= target {
			diff := j - i + 1
			if minLen > diff {
				minLen = diff
			}
			sum -= nums[i]
			i++
		}
		j++
	}

	return minLen % (len(nums) + 1)
}

