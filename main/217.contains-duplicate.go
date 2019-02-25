package main

import "fmt"

/**

题目:
给定一个整数数组，判断是否存在重复元素。
如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

示例 1:
输入: [1,2,3,1]
输出: true

示例 2:
输入: [1,2,3,4]
输出: false

示例 3:
输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
*/

func main() {
	nums := []int{1, 2, 3}
	r := containsDuplicate(nums)
	fmt.Println(r)
}

/**
方法1:
空间换时间，来一个map，如果已为1则直接返回false，否则就设置为1


时间复杂度：O(n)
空间复杂度：O(n)
*/
func containsDuplicate(nums []int) bool {
	tempMap := make(map[int]int)

	for _,v := range nums {
		if tempMap[v] > 0 {
			return true
		} else {
			tempMap[v] = 1
		}
	}
	return false
}
