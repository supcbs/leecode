package main

import "fmt"

/**

题目:

给定长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，
其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:
输入: [1,2,3,4]
输出: [24,12,8,6]
说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）

*/

func main() {
	nums := []int{1, 2, 3, 4}
	r := productExceptSelf(nums)
	fmt.Println(r)
}

/**
方法1:

想要在O(n)完成，那么只能遍历n的常数倍。
想象一下，每个数的乘积应该是等于其左边的数乘积 * 右边的数乘积

那么从左往右，再从右往左遍历一遍即可，自然算出了最后答案


时间复杂度：O(n)
空间复杂度：O(1)
*/
func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return []int{}
	}

	ret := make([]int, numsLen)
	left := 1
	right := 1

	for i:=0;i<numsLen;i++ {
		ret[i] = left
		left *= nums[i]
	}

	for i:=numsLen-1;i>=0;i-- {
		ret[i] *= right
		right *= nums[i]
	}

	return ret
}


/**
方法2:

想要在O(n)完成

来一个初始化每个数都是1的temp数组，记录各个idx位置的数。
---
这种方法复杂度不满足要求

时间复杂度：O(n)
空间复杂度：O(1)
*/
func productExceptSelf2(nums []int) []int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return []int{}
	}

	ret := make([]int, numsLen)
	for k, _ := range ret {
		ret[k] = 1
	}

	for numK, numV := range nums {
		for tempK, _ := range ret {
			if tempK != numK {
				ret[tempK] *= numV
			}
		}
	}

	return ret
}
