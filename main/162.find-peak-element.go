package main

import "fmt"

/**

题目:
峰值元素是指其值大于左右相邻值的元素。
给你一个输入数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞

示例 1：
输入：nums = [1,2,3,1]
输出：2
解释：3 是峰值元素，你的函数应该返回其索引 2。
示例 2：
输入：nums = [1,2,1,3,5,6,4]
输出：1 或 5
解释：你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
 
提示：
1 <= nums.length <= 1000
-231 <= nums[i] <= 231 - 1
对于所有有效的 i 都有 nums[i] != nums[i + 1]
 
进阶：你可以实现时间复杂度为 O(logN) 的解决方案吗？
*/

func main() {
	nums := []int{6,5,4,3,2,3,2}
	r := findPeakElement(nums)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func findPeakElement(nums []int) int {
	// 边界条件判断
	if len(nums) <= 0 {
		return -1
	}

	// 初始值设置
	l := len(nums)
	left := 0
	right := len(nums) - 1
	mid := -1

	// 二分法判断
	for left <= right {
		mid = left + (right - left) / 2
		if mid - 1 >= 0 && nums[mid] < nums[mid-1] {
			right = mid - 1
		} else if mid + 1 < l && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return mid
}

func findPeakElementOK(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] > nums[mid+1] {
			r = mid
		} else {
			l = mid+1
		}
	}
	return l
}