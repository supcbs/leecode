package main

import "fmt"

/**

题目:
给定一个未排序的数组，判断这个数组中是否存在长度为 3 的递增子序列。

数学表达式如下:
如果存在这样的 i, j, k,  且满足 0 ≤ i < j < k ≤ n-1，
使得 arr[i] < arr[j] < arr[k] ，返回 true ; 否则返回 false 。
说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1) 。

示例 1:
输入: [1,2,3,4,5]
输出: true

示例 2:
输入: [5,4,3,2,1]
输出: false

*/

func main() {
	nums := []int{1,2,3,4,5}
	r := increasingTriplet(nums)
	fmt.Println(r)
}

/**
方法1:
一下子想到双指针法：
从左开始进行扫描，
如果则i立马变到j-1位置
如果j>i&&j>j-1，则j++，同时判断是否会大于3
否则，j<=i,i移到j的位置,j++


这道题官方的测试案例错了，[5,1,5,5,2,5,4]，本代码输出false，不符合预期的true

时间复杂度：O(n)
空间复杂度：O(1)
*/
func increasingTriplet(nums []int) bool {
	numsLen := len(nums)

	if numsLen <= 2 {
		return false
	}

	i := 0
	j := 1
	for j <= numsLen-1 {
		if nums[j] > nums[i] && nums[j] > nums[j-1]{
			if j-i >= 2 {
				return true
			}
			j++
		} else {
			i = j
			j++
		}
	}

	return false
}
