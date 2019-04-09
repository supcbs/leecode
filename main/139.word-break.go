package main

import (
	"fmt"
	"sort"
)

/**

题目:
给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，
判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：
拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。

示例 1：
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。

示例 2：
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。

示例 3：
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

*/

func main() {
	str := "leetcode"
	wordDict := []string{"leet", "code"}
	r := wordBreak(str, wordDict)
	fmt.Println(r)
}

/**
方法1:

这道题思路：每个字符后面加上空格，空格前证明ok了，那么空格后的值如果也在字典中，则也是ok的

查询采用空间换时间，新增数组map
遍历位数的时候，采用跳跃式遍历，根据现存的字母长度来

时间复杂度：O(n)
空间复杂度：O(1)
*/
func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) <= 0 {
		return false
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		dict[w] = true
		length[len(w)] = true
	}

	sizes := make([]int,0)
	for l := range length {
		sizes = append(sizes, l)
	}

	sort.Ints(sizes)

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		// 说明分割前半部分不ok，也没必要继续验证后面的了
		if dp[i] == false {
			continue
		}

		// 根据词汇长度进行遍历
		for _, size := range sizes {
			if i+size <= len(s) {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}
	}
	return dp[len(s)]
}
