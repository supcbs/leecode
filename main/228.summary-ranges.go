package main

import (
	"fmt"
	"strconv"
)

/**

题目:
给定一个无重复元素的有序整数数组，返回数组区间范围的汇总。

示例 1:
输入: [0,1,2,4,5,7]
输出: ["0->2","4->5","7"]
解释: 0,1,2 可组成一个连续的区间; 4,5 可组成一个连续的区间。

示例 2:
输入: [0,2,3,4,6,8,9]
输出: ["0","2->4","6","8->9"]
解释: 2,3,4 可组成一个连续的区间; 8,9 可组成一个连续的区间。

*/

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	nums = []int{0, 2, 3, 4, 6, 8, 9}
	r := summaryRangesOK(nums)
	fmt.Println(r)
}

/**
方法1:

这道题有类似的，叫做合并区间，给忘记题号了

思路大概就是从左往右遍历，检查是否存在n+1的数，存在的话继续找end的数，不存在就直接返回范围了

时间复杂度：O(n)
空间复杂度：O(1)
*/
func summaryRanges(nums []int) []string {
	numsLen := len(nums)
	if numsLen <= 0 {
		return []string{}
	}

	var ret []string
	startIdx := 0
	endIdx := 0
	for i := 0; i < numsLen-1; i++ {
		if i+1 <= numsLen && nums[i+1] != nums[i]+1 {
			if startIdx == endIdx {
				ret = append(ret, strconv.Itoa(nums[i]))
			} else {
				ret = append(ret, strconv.Itoa(nums[startIdx])+"->"+strconv.Itoa(nums[endIdx]))
			}

			startIdx = i + 1
			endIdx = i + 1
			continue
		}

		endIdx++
	}

	// 最后以为判断
	if startIdx == endIdx {
		ret = append(ret, strconv.Itoa(nums[endIdx]))
	} else {
		ret = append(ret, strconv.Itoa(nums[startIdx])+"->"+strconv.Itoa(nums[endIdx]))
	}

	return ret
}


/**
方法2:


还有一种方式类似于桶排序，因为排序是好的，所以只是分一下范围桶

时间复杂度：O(n)
空间复杂度：O(1)
*/
func summaryRangesOK(nums []int) []string {
	numsLen := len(nums)
	if numsLen <= 0 {
		return []string{}
	}

	bucket := make([][]int, 0)
	tmp := make([]int, 1)
	tmp[0] = nums[0]
	for i:=1 ; i < numsLen;i++{
		if nums[i] == nums[i-1] + 1 {
			tmp = append(tmp, nums[i])
		} else {
			bucket = append(bucket, tmp)
			tmp = make([]int, 1)
			tmp[0] = nums[i]
		}
	}

	if len(tmp) != 0 {
		bucket = append(bucket, tmp)
	}

	ret := make([]string, len(bucket))
	for k,v := range bucket {
		if len(v) == 1 {
			ret[k] = strconv.Itoa(v[0])
		} else if len(v) > 1 {
			ret[k] = strconv.Itoa(v[0]) +"->"+ strconv.Itoa(v[len(v)-1])
		}
	}

	return ret
}
