package main

import (
	"fmt"
)

/**

题目:
给定一个整数数组，判断数组中是否有两个不同的索引 i 和 j，
使得 nums [i] 和 nums [j] 的差的绝对值最大为 t，
并且 i 和 j 之间的差的绝对值最大为 ķ。

示例 1:
输入: nums = [1,2,3,1], k = 3, t = 0
输出: true

示例 2:
输入: nums = [1,0,1,1], k = 1, t = 2
输出: true

示例 3:
输入: nums = [1,5,9,1,5,9], k = 2, t = 3
输出: false
*/

func main() {
	nums := []int{2, 2}
	k := 3
	t := 0
	//r := containsNearbyAlmostDuplicate(nums, k, t)
	r := containsNearbyAlmostDuplicateOK(nums, k, t)

	fmt.Println(r)
}

/**
方法1:
由于ij之间有最大距离的限制，所以需要采用滑动窗口 + 定点遍历窗口的方式进行处理

确定滑动窗口的范围：i和j。i一直递增即可，j与i的距离不超过k位，即j=i-k>0的情况
然后遍历>=j到<i的位置进行判断绝对值是否满足要求。


时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 0 || t < 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {

		// 滑动窗口范围
		j := 0
		if i-k > 0 {
			j = i - k
		}

		// 滑动窗口内遍历
		for ; j >= 0 && j < i; j++ {
			if nums[i]-nums[j] <= t && nums[j]-nums[i] <= t {
				return true
			}
		}
	}

	return false
}

/**
方法2:
时间换空间是经典的方式，此处依旧适用。
遍历每个数，将当前索引记录下来，map的key为num，value为idx索引值。
然后根据t与nums长度来决定使用的遍历方式。

如果t > len(nums),那说明都满足，直接进行遍历map看一下是否存在，当前num
如果t < len(nums),那么全部遍历会比较多，就直接寻找num上下间隔t的就好，然后在判断idx

这里为什么是t和nums的长度做比较，因为t涉及到两个数范围问题，如果说这个比较小，那肯定遍历num+-t比较短，否则遍历data比较短



时间复杂度：O(n)
空间复杂度：O(n)
*/
func containsNearbyAlmostDuplicateOK(nums []int, k int, t int) bool {
	if k < 0 || t < 0 {
		return false
	}

	data := make(map[int]int)
	for idx, num := range nums {
		if len(nums) < t {
			for dataK, dataV := range data {
				if idx-dataV <= k && dataV-idx <= k {
					if dataK-num <= t && num-dataK <= t {
						return true
					}
				}
			}
		} else {
			for searchNum := num - t; searchNum <= num+t; searchNum++ {
				if _, exist := data[searchNum]; !exist {
					continue
				}

				if idx-data[searchNum] <= k {
					return true
				}
				if searchNum == num {
					data[num] = idx
				}
			}
		}

		data[num] = idx
	}

	return false
}
