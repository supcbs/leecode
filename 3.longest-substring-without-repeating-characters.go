package main

import (
	"fmt"
)

/**

题目:

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: "pwwkew" abba dwdf
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

*/

func main() {
	s := "abba"
	l := lengthOfLongestSubstring(s)
	fmt.Println("result",l)
}

/**
	滑动窗口判断
 */
func lengthOfLongestSubstring(s string) int {
	var i, j, l int
	var currString string
	temp := make(map[string]int)
	sLen := len(s)

	for ; j < sLen; j++ {
		currString = s[j : j+1]
		fmt.Println(currString)
		if _, ok := temp[currString]; ok {
			offset := temp[currString]
			if offset > i {
				i = offset
			}
			temp[currString] = j
		}

		if j-i+1 > l {
			l = j - i + 1
		}
		temp[currString] = j + 1
	}

	return l
}
