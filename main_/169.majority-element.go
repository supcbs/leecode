package main

import "fmt"

/**

题目:
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:
输入: [3,2,3]
输出: 3

示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
*/

func main() {
	nums := []int{2,2,1,1,1,2,2}
	r := majorityElementNew(nums)
	fmt.Println(r)
}

func majorityElementNew(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0
	}

	targetNum := nums[0]
	acc := 1
	for i:=1;i<numsLen;i++ {
		if acc == 0 {
			targetNum = nums[i]
		}
		if nums[i] == targetNum {
			acc++
		} else {
			acc--
		}
	}

	return targetNum
}

/**
方法1:
从第一个数开始遍历，遇到相同的就+1，不同的就-1，当count为0的时候，则指定当前遍历到的数为众数
---

这个算法有个专业名词：Boyer-Moore 投票算法 （摩尔投票算法）


时间复杂度：O(n)
空间复杂度：O(1)
*/
func majorityElement(nums []int) int {
	var targetNum,count int

	for _,v :=range nums {
		if count == 0 {
			targetNum = v
		}

		if targetNum == v {
			count ++
		} else {
			count --
		}
	}

	return targetNum
}
