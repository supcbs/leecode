package main

import (
	"fmt"
)

/**

题目:
给你字符串 s 和整数 k 。
请返回字符串 s 中长度为 k 的单个子字符串中可能包含的最大元音字母数。
英文中的 元音字母 为（a, e, i, o, u）。

示例 1：
输入：s = "abciiidef", k = 3
输出：3
解释：子字符串 "iii" 包含 3 个元音字母。
示例 2：
输入：s = "aeiou", k = 2
输出：2
解释：任意长度为 2 的子字符串都包含 2 个元音字母。

*/



func main() {
	s := "abciiidef"
	r := maxVowels(s, 3)
	fmt.Println(r)
}

/**
方法1:

没限制人员的解法


时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxVowels(s string, k int) int {
	// 边界条件判断
	if len(s) < 0 || k == 0 {
		return 0
	}

	// 初始化必要参数
	i, j := 0, 0
	sum, max := 0, 0
	data := make(map[string]int, 0)
	data["a"] = 1
	data["e"] = 1
	data["i"] = 1
	data["o"] = 1
	data["u"] = 1

	// 滑动窗口处理。
	for j < len(s) {
		for j - i + 1 > k {
			if _, ok := data[string(s[i])]; ok {
				sum--
			}
			i++
		}
		if _, ok := data[string(s[j])]; ok {
			sum++
		}

		if sum > max {
			max = sum
		}
		j++
	}

	return max
}