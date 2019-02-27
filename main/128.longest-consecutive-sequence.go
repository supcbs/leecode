package main

import "fmt"

/**

题目:
给定一个未排序的整数数组，找出最长连续序列的长度。
要求算法的时间复杂度为 O(n)。

示例:
输入: [100, 4, 200, 1, 3, 2]
输出: 4

解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

*/

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	r := longestConsecutive(nums)
	fmt.Println(r)
}

/**
方法1:

要求了时间复杂度，那么很容易想到拿一个map来存一下。
然后再遍历nums，若前一个数没有，则这个数可以是起始的数s，然后一直往后遍历直到map中不存在e
那么此时的e-s就是最大的长度

时间复杂度：O(n)
空间复杂度：O(n)
*/
func longestConsecutive(nums []int) int {
	// 不能用slice，最大的数是数组长度，因为那样太长了。
	numsMap := make(map[int]bool)

	for _, v := range nums {
		numsMap[v] = true
	}

	maxLen := 0
	for _, v := range nums {
		_, ok := numsMap[v-1]

		// 存在则不是起始的数
		if ok {
			continue
		}

		t := v + 1
		for _, ok := numsMap[t]; ok; _, ok = numsMap[t] {
			t++
		}

		if t-v > maxLen {
			maxLen = t - v
		}
	}
	return maxLen
}
