package main

import (
	"fmt"
)

/**

题目:
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。

进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

*/

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	r := maxSubArray(nums)
	fmt.Println(r)
}

/**
方法1:

直接遍历数组进行累加。
如果累加数小于0，则重新算
如果大于0，则继续累加起来

同时暂存最大的数

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxSubArray(nums []int) int {
	numsLen := len(nums)

	if numsLen <= 0 {
		return 0
	}

	ret := nums[0]
	sumNum := 0
	for _, v := range nums {
		if sumNum > 0 {
			sumNum += v
		} else {
			sumNum = v
		}
		if sumNum > ret {
			ret = sumNum
		}
	}

	return ret
}

/**
方法2:

这道题可以理解为一个动态规划的题目：

公式：f(n) = max(f(n-1) + A[n], A[n])

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxSubArrayOK(nums []int) int {
	numsLen := len(nums)

	if numsLen <= 0 {
		return 0
	}

	ret := nums[0]
	fn := 0
	for _, v := range nums {
		if fn + v > v {
			fn = fn + v
		} else {
			fn = v
		}

		if fn > ret {
			ret = fn
		}
	}

	return ret
}

func maxSubArray(nums []int) int {
	return maxSubArrayRecursion(nums, 0, len(nums) - 1).mSum
}

type Status struct {
	lSum, rSum, mSum, iSum int
}

// 既然是分治，先摆出递归的样子
func maxSubArrayRecursion(nums []int, l ,r int) Status {
	// 边界条件
	if l == r {
		return Status{nums[l],nums[l],nums[l],nums[l]}
	}

	// 区分两个区间
	mid := (l+r) >> 1
	// 左区间
	lSub := maxSubArrayRecursion(nums, l, mid)
	// 右区间
	rSub := maxSubArrayRecursion(nums, mid+1, r)

	return pushUp(lSub, rSub)
}

func pushUp(l, r Status) Status {
	// 区间和
	iSum := l.iSum + r.iSum
	// 以l为左端点最大
	lSum := Max(l.lSum, l.iSum + r.lSum)
	// 以r为左端点最大
	rSum := Max(r.rSum, r.iSum + l.rSum)
	// 以最大字段和
	mSum := Max(Max(l.mSum, r.mSum), l.rSum + r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func Max(i,j int) int {
	if i > j {
		return j
	}
	return i
}