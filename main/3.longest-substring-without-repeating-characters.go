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
	s := "abcabcbb"
	l := lengthOfLongestSubstring(s)
	fmt.Println("result",l)
}

/**
方法1:

和76、209类似
双指针+滑动窗口

ij指针从0开始往前，j每往前一步就将当前的字母记录到map中，key为字母，value位置+1
每次当有重复的字母出现的时候，并且记录的offset大于i，则i就变为map中记录的位置。
每次都计算长度，j-i+1就是长度


时间复杂度：O(n)
空间复杂度：O(1)
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) < 0 {
		return 0
	}

	var temp [128]int
	var i,j,maxLen int
	for ;j<len(s);j++{
		curStr := s[j]
		if temp[curStr] > i {
			i = temp[curStr]
		}

		if j-i+1 > maxLen {
			maxLen = j-i+1
		}
		temp[curStr] = j+1
	}

	return maxLen
	//if len(s) < 0 {
	//	return 0
	//}
	//
	//var i,j,maxLen int
	//temp := make(map[byte]int)
	//for ;j<len(s);j++{
	//	curStr := s[j]
	//	if offset,ok:=temp[curStr];ok{
	//		// 需要往后
	//		if offset > i{
	//			i = offset
	//		}
	//	}
	//
	//	if j-i+1>maxLen {
	//		maxLen = j-i+1
	//	}
	//	temp[curStr] = j+1
	//}
	//
	//return maxLen
}

/**
方法2:

最优的考虑了两个点：
1.当存在重复位置的时候，是不不需要判断是否最长的
2.当temp中的offset大于i的位置的时候，说明重复了，直接将i进行换位即可

时间复杂度：O(n)
空间复杂度：O(1)
*/
func lengthOfLongestSubstringOK(s string) int {
	if len(s) < 0 {
		return 0
	}

	var i,j,maxLen int
	temp := [128]int{}
	for ;j<len(s);j++{
		curStr := s[j]
		if temp[curStr] > i {
			i = temp[curStr]
		} else if j-i+1>maxLen {
			maxLen = j-i+1
		}
		temp[curStr] = j+1
	}

	return maxLen
}
