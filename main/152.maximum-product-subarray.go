package main

import "fmt"

/**

题目:
给定一个整数数组 nums ，找出一个序列中乘积最大的连续子序列（该序列至少包含一个数）。

示例 1:
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。

示例 2:
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。

*/

func main() {
	nums := []int{-2,0,-1}
	nums = []int{0,2}
	nums = []int{1,2,-1,-2,2,1,-2,1,4,-5,4}
	r := maxProduct(nums)
	fmt.Println(r)
}

/**
方法1:

这道题雷同53题，以一个是加起来最长，一个是乘起来最大

f(n) = max(f(n-1) * A[n], A[n])
但是这把有个细节，就是min遇到负数就是最大的了，max遇到负数就是最小的了

所以为了体现这两个的这两个值需要多一个变量

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxProduct(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 0{
		return 0
	}

	if numsLen == 1 {
		return nums[0]
	}

	maxFn := 1
	minFn := 1
	ret := ^int(uint(0) >> 1)
	for i:=0;i<numsLen;i++{

		// 如果当前数是负数就对换，为后面得到的数依旧是max和min
		if nums[i] < 0 {
			maxFn, minFn = minFn, maxFn
		}
		tempFn := maxFn * nums[i]
		if tempFn >= nums[i] {
			maxFn = tempFn
		} else {
			maxFn = nums[i]
		}

		tempFn = minFn * nums[i]
		if tempFn <= nums[i] {
			minFn = tempFn
		} else {
			minFn = nums[i]
		}
		fmt.Println(maxFn, minFn)



		if maxFn > ret {
			ret = maxFn
		}
	}

	return ret
}
