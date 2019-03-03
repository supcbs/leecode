package main

import (
	"fmt"
	"sort"
)

/**

题目:
给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

示例 1:
输入: nums = [1, 5, 1, 1, 6, 4]
输出: 一个可能的答案是 [1, 4, 1, 5, 1, 6]

示例 2:
输入: nums = [1, 3, 2, 2, 3, 1]
输出: 一个可能的答案是 [2, 3, 1, 3, 1, 2]

说明:
你可以假设所有输入都会得到有效的结果。

进阶:
你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？



*/

func main() {
	nums := []int{4, 3, 2, 1, 5, 6}
	wiggleSort(nums)
	fmt.Println(nums)
}

/**
方法1:

先排序，后分成两个数组，其中一个顺序穿插，另一个倒序穿插


时间复杂度：O(n)
空间复杂度：O(1)
*/
func wiggleSort(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}

	// 不能直接tmp = nums,因为是指针引用所以会又问题
	tmp := make([]int, numsLen)
	for k, v := range nums {
		tmp[k] = v
	}
	sort.Ints(tmp)
	j := numsLen - 1
	//先奇数
	for i:=1;i<len(nums);i,j=i+2,j-1{
		nums[i] = tmp[j]
	}
	//后偶数
	for i:=0;i<len(nums);i,j=i+2,j-1{
		nums[i] = tmp[j]
	}
	return
}
