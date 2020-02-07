package main

import "fmt"

/**

题目:
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。

示例 1:
输入: [2,3,1,1,4]
输出: true
解释: 从位置 0 到 1 跳 1 步, 然后跳 3 步到达最后一个位置。

示例 2:
输入: [3,2,1,0,4]
输出: false

解释: 无论怎样，你总会到达索引为 3 的位置。但该位置的最大跳跃长度是 0 ， 所以你永远不可能到达最后一个位置。
*/

func main() {
	nums := []int{2, 0, 0}
	r := canJump(nums)
	fmt.Println(r)
}

/**
方法1:

倒着数，也就是l-1的数加上当前位置必须要>=下一个位置，才可以。
如果最后一个都可以，则位置到了0的位置

时间复杂度：O(n)
空间复杂度：O(1)
*/
func canJump(nums []int) bool {

	numsLen := len(nums)
	// 这个就是下一个位置
	k := numsLen - 1
	for i := numsLen - 2; i >= 0; i-- {
		if i+nums[i] >= k {
			k = i
		}
	}

	return k == 0
}

/**
方法2:

正着数，设置当前max = max(nums[i] + 1, max),如果i>max则说明在这一步无论如何也走不到最后

时间复杂度：O(n)
空间复杂度：O(1)
*/
func canJump2(nums []int) bool {

	numsLen := len(nums)
	max := 0
	for i := 0; i < numsLen; i++ {
		if i > max {
			return false
		}

		if nums[i]+i > max {
			max = nums[i] + i
		}
	}

	return true
}
