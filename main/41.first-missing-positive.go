package main

import "fmt"

/**

题目:
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:
输入: [1,2,0]
输出: 3

示例 2:
输入: [3,4,-1,1]
输出: 2

示例 3:
输入: [7,8,9,11,12]
输出: 1

说明:你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。
*/

func main() {
	nums := []int{1, 2, 3}
	r := firstMissingPositive(nums)
	fmt.Println(r)
}

/**
方法1:
可以理解为未出现过的正整数（>0）一定出现在两种场景，假设数组只有5位：
a.当前5位存在非12345的时候，则缺哪位就说明那位就是未出现的最小的正整数
b.当前5位正好是12345的时候，则最小正整数就是n+1=6

所以遍历两次，第一次标记前n位小数，第二次遍历判断是否存在前n位小数

时间复杂度：O(n)
空间复杂度：O(1)
*/
func firstMissingPositive(nums []int) int {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 1
	}

	// 这一步会初始化为0,这里至少多1位因为首位是0
	temp := make([]int, numsLen+1)

	for _,num := range nums {
		if num > numsLen || num <= 0 {
			continue
		}

		temp[num] = 1
	}

	for j := 1; j <= numsLen; j++ {
		if temp[j] == 0 {
			return j
		}
	}

	return numsLen + 1
}
