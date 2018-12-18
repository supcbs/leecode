package main

import "fmt"

/**

题目:

给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9

所以返回 [0, 1]

*/

func main() {
	nums := []int{2, 7, 11, 5}
	target := 9

	r := twoSum(nums, target)
	fmt.Println(r)
}

// 方法1：
// 算法复杂度：n^2 空间复杂度：1
func twoSum1(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 0 {
		return nil
	}

	returnArr := make([]int, 2)
sum:
	for i := 0; i < numsLen; i++ {
		for j := i + 1; j < numsLen; j++ {
			if nums[i]+nums[j] == target {
				returnArr[0] = i
				returnArr[1] = j
				break sum
			}
		}
	}

	return returnArr
}

// 方法2：
// 算法复杂度：n 空间复杂度：n
func twoSum2(nums []int, target int) []int {
	cacheMap := make(map[int]int, 0)

	for k, v := range nums {
		cacheMap[v] = k
	}

	numsLen := len(nums)
	returnArr := make([]int, 2)
	for i := 0; i < numsLen; i++ {
		needNum := target - nums[i]

		if j, ok := cacheMap[needNum]; ok {
			returnArr[0] = i
			returnArr[1] = j
		}
	}

	return returnArr
}

// 方法3：
// 算法复杂度：n 空间复杂度：n
func twoSum(nums []int, target int) []int {
	cacheMap := make(map[int]int, 0)
	numsLen := len(nums)
	returnArr := make([]int, 2)
	for i:=0; i<numsLen;i++{
		needNum := target - nums[i]

		if j,ok := cacheMap[needNum]; ok {
			returnArr[0] = i
			returnArr[1] = j
			break
		}

		cacheMap[nums[i]]=i
	}
	return returnArr
}
