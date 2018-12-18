package main

import "fmt"

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
*/

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	sLen := len(s)
	var i, j, maxLen int
	cache := make(map[string]int)
	var returnStr, currStr string

	for ; j < sLen; j++ {
		currStr = s[j : j+1]

		if _, ok := cache[currStr]; ok {
			i = cache[currStr]
		}

		if j - i + 1 > maxLen {
			maxLen = j - i + 1
			returnStr = s[i:j+1]
		}

		cache[currStr] = j + 1
	}

	return returnStr
}
