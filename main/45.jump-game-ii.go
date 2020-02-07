package main

import "fmt"

/**

题目:
给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例:
输入: [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
*/

func main() {
	nums := []int{2, 1}
	r := jumpOK(nums)
	fmt.Println(r)
}

/**
方法1：

贪心算法：选择下一步最远的点。

大概描述：
1.标记最大的位置，当i==最大位置的时候，则认为步数加1。
2.当最大位置大于等于数组长度的时候，则认为停止了

时间复杂度：O(n)
空间复杂度：O(1)
 */

func jumpOK(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 1 {
		return 0
	}

	var maxPos, end, step int
	for i := 0; i < numsLen; i++ {
		if nums[i]+i > maxPos {
			maxPos = nums[i] + i
		}

		if maxPos > numsLen {
			break
		}

		if i == end {
			end = maxPos
			step++
		}
	}

	return step
}

/**
方法2:

贪心算法：
最简单的贪心算法是每次都选择最远的点。
但是这道题不能这么选择，需要选择下一步最远的点。这样就没问题了

时间复杂度：O(n)
空间复杂度：O(1)
*/
func jump1(nums []int) int {
	numsLen := len(nums)
	var i, step int

	for i < numsLen-1 {
		max := 0
		jMax := i + nums[i]
		for j := i + 1; j <= jMax && j < numsLen; j++ {
			if j == numsLen-1 {
				i = j
				break
			}

			canJump := j + nums[j]
			if canJump > max {
				max = canJump
				i = j
			}
		}

		step++
	}

	return step
}

/**
方法3:

动态规划：
算出每个位置最短的距离，在此基础上算下一个距离。

这个会超时。。。。

时间复杂度：O(n)
空间复杂度：O(n)
*/
func jump3(nums []int) int {
	numsLen := len(nums)

	data := make([]int, numsLen)

	for i, num := range nums {
		minLen := data[i]
		for j := i + 1; j <= num+i && j < numsLen; j++ {
			if data[j] == 0 {
				data[j] = minLen + 1
			}
		}
	}

	return data[numsLen-1]
}