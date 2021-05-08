package main

import "fmt"

/**

题目:
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。

进阶：
你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
 
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
 

提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109

*/

func main() {
	nums := []int{5}
	r := searchRange(nums, 5)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func searchRange(nums []int, target int) []int {
	start := -1
	end := -1
	// 边界值判断
	if len(nums) <= 0 {
		return []int{start, end}
	}

	// 初始值准备
	left := 0
	right := len(nums) - 1
	hit := -1

	// 二分查找
	for left <= right {
		mid := left + (right - left) / 2
		if target == nums[mid] {
			hit = mid
			break
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}

	// 确定边界
	if hit < 0 {
		return []int{start, end}
	}

	// 确定开始位置
	start = hit
	for start >= 0 && target == nums[start] {
		start--
	}
	start++

	// 确定右边界
	end = hit
	for end <= len(nums)-1 && target == nums[end] {
		end++
	}
	end--

	return []int{start, end}
}