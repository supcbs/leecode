package main

import "fmt"

/**
题目：

给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
nums1 = [1, 3]
nums2 = [2]
则中位数是 2.0

示例 2:
nums1 = [1, 2]
nums2 = [3, 4]
则中位数是 (2 + 3)/2 = 2.5

*/

func main() {
	nums1 := []int{1,3}
	nums2 := []int{2}

	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)
	fmt.Println(5/2)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {


	return 0
}
