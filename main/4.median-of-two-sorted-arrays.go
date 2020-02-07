package main

import (
	"fmt"
	"sort"
)

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
	nums1 := []int{1, 3}
	nums2 := []int{2}

	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)
}

/**
方法1:
双指针法：
算出总长度len，因为golang是向下取整，k=（len+1）/2得到的肯定是中位数的位置。
那么开始遍历0->k的位置，指定两个指针分别指向两个数组，
a.当指针都小于各自长度的的时候，比较指针所指的数，谁小谁为目标mid，并且指针右移。
b.当其中一个数组走到尽头，则另一个没走到尽头的数组，移动剩余的位置，即为目标mid

如果此时len为奇数，则上文的数就是目标数
为偶数的话，需要根据当前指针的位置计算下一个最小的目标数。


这道题真痛苦：其他解法：https://blog.csdn.net/hang404/article/details/84786904

时间复杂度：O(logn)
空间复杂度：O(1)
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2 = len(nums1), len(nums2)
	var stop = (l1 + l2 + 1) / 2
	var p1, p2 = 0, 0
	var m1 = 0
	for i := 0; i < stop; i++ {
		if p1 < l1 && p2 < l2 {
			if nums1[p1] < nums2[p2] {
				m1 = nums1[p1]
				p1++
			} else {
				m1 = nums2[p2]
				p2++
			}
		} else if p1 < l1 {
			m1 = nums1[p1]
			p1++
		} else if p2 < l2 {
			m1 = nums2[p2]
			p2++
		}
	}
	if (l1+l2)%2 != 0 {
		return float64(m1)
	}

	var m2 = 0
	if p1 < l1 && p2 < l2 {
		if nums1[p1] < nums2[p2] {
			m2 = nums1[p1]
		} else {
			m2 = nums2[p2]
		}
	} else if p1 < l1 {
		m2 = nums1[p1]
	} else if p2 < l2 {
		m2 = nums2[p2]
	}
	return float64(m1+m2) * 0.5
}


func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	// 保证nums1长度小于nums2
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 特例1: 包含了不能再减少区间的特殊情况(并不都是特殊的, 但数量少无所谓)
	if len(nums1) == 0 || len(nums2) <= 4 {
		if len(nums1) > 0 {
			for _, num := range nums1 {
				nums2 = append(nums2, num)
			}
			sort.Ints(nums2)
		}
		if len(nums2) % 2 == 1 {
			return float64(nums2[len(nums2)/2])
		} else {
			return float64(nums2[len(nums2)/2] + nums2[len(nums2)/2-1])/2
		}
	}
	// 特例2: 有一个区间长度为1, 另一个长度大于3的情况
	if len(nums1) == 1 && len(nums2) > 3 {
		if len(nums2) % 2 == 1 {
			nums2 = nums2[len(nums2)/2-1:len(nums2)/2+2] //保留3位
		} else {
			nums2 = nums2[len(nums2)/2-1:len(nums2)/2+1] //保留2位
		}
		return findMedianSortedArrays(nums1, nums2)
	}
	// 特例3: 有一个区间长度为2, 另一个区间大于4
	if len(nums1) == 2 && len(nums2) > 4 {
		if len(nums2) % 2 == 1 {
			nums2 = nums2[len(nums2)/2-1:len(nums2)/2+2] // 保留3位
		} else {
			nums2 = nums2[len(nums2)/2-2:len(nums2)/2+2] // 保留4位
		}
		return findMedianSortedArrays(nums1, nums2)
	}
	// 能减少区间的情况
	odd1 := len(nums1) % 2
	odd2 := len(nums2) % 2
	mid1 := 0.0
	mid2 := 0.0
	if odd1 == 1 {
		mid1 = float64(nums1[len(nums1)/2])
	} else {
		mid1 = float64(nums1[len(nums1)/2] + nums1[len(nums1)/2-1])/2
	}
	if odd2 == 1 {
		mid2 = float64(nums2[len(nums2)/2])
	} else {
		mid2 = float64(nums2[len(nums2)/2] + nums2[len(nums2)/2-1])/2
	}

	// 如果恰好相等, 直接返回
	if mid1==mid2 {
		return mid1
	}

	// 两边能减少的最大的数量
	removeNum := 0
	if odd1 == 1 {
		removeNum = len(nums1)/2
	} else {
		removeNum = len(nums1)/2-1
	}
	if odd2 == 1{
		if len(nums2)/2 < removeNum {
			removeNum = len(nums2)/2
		}
	} else {
		if len(nums2)/2-1 < removeNum {
			removeNum = len(nums2)/2-1
		}
	}

	if mid2 > mid1 {
		nums1 = nums1[removeNum:]
		nums2 = nums2[0:len(nums2)-removeNum]
	} else {
		nums1 = nums1[0:len(nums1)-removeNum]
		nums2 = nums2[removeNum:]

	}
	return findMedianSortedArrays(nums1, nums2)
}
