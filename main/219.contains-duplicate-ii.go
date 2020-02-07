package main

import "fmt"

/**

题目:
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，
使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

示例 1:
输入: nums = [1,2,3,1], k = 3
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1
输出: true

示例 3:
输入: nums = [1,2,3,1,2,3], k = 2
输出: false
*/

func main() {
	nums := []int{1,2,3,1,2,3}
	r := containsNearbyDuplicate(nums, 2)
	fmt.Println(r)
}

/**
方法1:
空间换时间，来一个map，
若当前值的map对应的value不为0，并且i减去map对应的value小于k
否则将当前值对应的map的value置为i+1


时间复杂度：O(n)
空间复杂度：O(1)
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if k < 0 {
		return false
	}

	m := make(map[int]int, len(nums))
	for i, n := range nums {
		if m[n] != 0 && i - m[n] + 1 <= k {
			return true
		}

		// 这里+1是因为，当首位的时候，也为0，需要区分开来，所以加一个1，在判断小于k的时候也需要加上
		m[n] = i + 1
	}

	return false
}


func containsNearbyDuplicateOK(nums []int, k int) bool {
	numsLen := len(nums)
	if numsLen <= 0 || k <= 0 {
		return false
	}

	// key记录具体数字。value记录索引
	tmpMap := make(map[int]int)
	for index,num := range nums {
		if _,ok := tmpMap[num]; ok && index - tmpMap[num] <= k {
			return true
		}

		tmpMap[num] = index
	}

	return false
}
