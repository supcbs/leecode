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
	nums := []int{-1,-1}
	k := 1
	t := -1
	//r := containsNearbyAlmostDuplicate(nums, k, t)
	r := containsNearbyAlmostDuplicate88(nums, k, t)

	fmt.Println(r)
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	numsLen := len(nums)
	if numsLen <= 0 {
		return false
	}
	w := t+1
	bucket := make(map[int]int)
	for i:=0;i<numsLen;i++ {
		index := getBucket(nums[i],w)
		// 如果当前桶存在则直接返回
		if _,ok := bucket[index];ok {
			return true
		}

		// 查看一下前一个桶是否存在
		if preNum,ok := bucket[index-1];ok {
			if nums[i] - preNum <= t {
				return true
			}
		}

		// 查看后一个桶是否存在
		if nextNum,ok := bucket[index+1];ok {
			if nextNum - nums[i] <= t {
				return true
			}
		}

		bucket[index] = nums[i]


		// 删除k之外的数，来保持桶当中的数都是符合|i-j| <= k
		if len(bucket)>k {
			delete(bucket,getBucket(nums[i-k],w))
		}
	}

	return false
}

func getBucket(num, w int) int {
	if num < 0 {
		return (num + 1) / w - 1
	} else {
		return num / w
	}
}

// 这里利用桶排序的思想, 把数据划分为[-t-1, -1], [0, t], [t+1, 2t+1], [2t+2, 3t+2]...
// 同一个桶内的元素间距一定小于等于t, 相邻桶的元素判断下间距即可
// 另外用map来做滑动窗口, 窗口的大小最大为k
func containsNearbyAlmostDuplicate88(nums []int, k int, t int) bool {
	if t < 0 || k <= 0 {
		return false
	}

	bucketRecord := make(map[int]int)
	var w int = t + 1

	for index, num := range nums {
		bucket := getBucket(num , w)
		_, has := bucketRecord[bucket]
		if has {
			return true
		}

		prevBucketNum, prevBucketHas := bucketRecord[bucket - 1]
		if prevBucketHas {
			if num >= prevBucketNum && num - prevBucketNum <= t {
				return true
			}
			if num < prevBucketNum && prevBucketNum - num <= t {
				return true
			}
		}
		nextBucketNum, nextBucketHas := bucketRecord[bucket + 1]
		if nextBucketHas {
			if num >= nextBucketNum && num - nextBucketNum <= t {
				return true
			}
			if num < nextBucketNum && nextBucketNum - num <= t {
				return true
			}
		}

		bucketRecord[bucket] = num
		// 窗口大小判断, 否则最早入窗口的出窗口
		if len(bucketRecord) > k {
			delete(bucketRecord, getBucket(nums[index - k], w))
		}
	}
	return false
}


/**
方法1:
由于ij之间有最大距离的限制，所以需要采用滑动窗口 + 定点遍历窗口的方式进行处理

确定滑动窗口的范围：i和j。i一直递增即可，j与i的距离不超过k位，即j=i-k>0的情况
然后遍历>=j到<i的位置进行判断绝对值是否满足要求。


时间复杂度：O(n^2)
空间复杂度：O(1)
*/
func containsNearbyAlmostDuplicate12(nums []int, k int, t int) bool {
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


func containsNearbyAlmostDuplicate99(nums []int, k int, t int) bool {
	numsLen := len(nums)
	if numsLen <= 0 {
		return false
	}
	w := t+1
	bucket := make(map[int]int)
	for i:=0;i<numsLen;i++ {
		index := nums[i] / w
		// 如果当前桶存在则直接返回
		if _,ok := bucket[index];ok {
			return true
		}

		// 查看一下前一个桶是否存在
		if _,ok := bucket[index-1];ok {
			if nums[i] - bucket[index-1] <= t {
				return true
			}
		}

		// 查看后一个桶是否存在
		if _,ok := bucket[index+1];ok {
			if bucket[index+1] - nums[i] <= t {
				return true
			}
		}

		bucket[index] = nums[i]


		// 删除k之外的数，来保持桶当中的数都是符合|i-j| <= k
		if i>k {
			if _, ok:= bucket[nums[i-k]/w];ok{
				delete(bucket,nums[i-k]/w)
			}
		}
	}

	return false
}



