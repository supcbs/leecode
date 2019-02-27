package main

import (
	"fmt"
)

/**

题目:
给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
如果数组元素个数小于 2，则返回 0。

示例 1:
输入: [3,6,9,1]
输出: 3
解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。

示例 2:
输入: [10]
输出: 0
解释: 数组元素个数小于 2，因此返回 0。
说明:

你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。

*/

func main() {
	nums := []int{1,1,1,1,1,5,5,5,5,5}
	r := maximumGap(nums)
	fmt.Println(r)
}

/**
方法1:

这道题的难点在于：线性时间复杂度和空间复杂度的条件下解决此问题
所以一开始想到直接快排+遍历求间隔的最大值的方式来解决，就不行了，因为快排就一会Nlogn了

这道题用到了一个叫做桶排序的东西。
需要构造最大桶，最小桶。首先算出数组中的最大数和最小数的差值，然后除以n-1向上取整，计算出桶之间的差值gap。
然后根据最大最小值差值除以数组长度，求出桶所需的长度
桶的key：(num - minNum)/len

然后遍历桶，前一个桶的最大值与后一个桶的最小值的差值，求出最大值就是最大差值了

桶排序：https://www.cnblogs.com/bywallance/p/5761269.html

时间复杂度：O(n)
空间复杂度：O(2n)
*/
func maximumGap(nums []int) int {
	numsLen := len(nums)
	if numsLen < 2 {
		return 0
	}

	// 计算最大最小值
	minNum := nums[0]
	maxNum := nums[0]
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
		if v < minNum {
			minNum = v
		}
	}

	if minNum == maxNum {
		return 0
	}

	if numsLen == 2 {
		return maxNum - minNum
	}

	// 计算桶间隔
	aveGap := (maxNum - minNum) / (numsLen - 1)
	if aveGap <= 0 {
		aveGap = 1
	}
	// 计算桶的个数
	bucketLen := (maxNum - minNum) / aveGap

	minBucket := make([]int, bucketLen+1)
	maxBucket := make([]int, bucketLen+1)

	const NUM_MAX = int(^uint(0) >> 1)
	const NUM_MIN = ^NUM_MAX
	for i := 0; i <= bucketLen; i++ {
		minBucket[i] = NUM_MAX
		maxBucket[i] = NUM_MIN
	}

	// 开始分桶
	for _, v := range nums {
		bucketIndex := (v - minNum) / aveGap
		if v < minBucket[bucketIndex] {
			minBucket[bucketIndex] = v
		}
		if v > maxBucket[bucketIndex] {
			maxBucket[bucketIndex] = v
		}
	}

	// 遍历桶求最大的数
	maxGap := 0
	preMaxBucketNum := maxBucket[0]

	for i := 1; i <= bucketLen; i++ {
		if maxBucket[i] == NUM_MIN && minBucket[i] == NUM_MAX {
			continue
		}

		tempGap := minBucket[i] - preMaxBucketNum
		if tempGap > maxGap {
			maxGap = tempGap
		}
		preMaxBucketNum = maxBucket[i]
	}

	return maxGap
}
