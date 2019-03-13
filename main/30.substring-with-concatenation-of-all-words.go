package main

import "fmt"

/**

题目:
给定一个字符串 s 和一些长度相同的单词 words。
找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
注意子串要与 words 中的单词完全匹配，中间不能有其他字符，
但不需要考虑 words 中单词串联的顺序。

示例 1：
输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoor" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。

示例 2：
输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]

*/

func main() {
	str := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	r := findSubstring(str, words)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func findSubstring(s string, words []string) []int {
	var ret []int
	if len(words) == 0 {
		return ret
	}

	// 将词汇放到map中
	wordMap := make(map[string]int)
	for _, v := range words {
		wordMap[v]++
	}

	// 计算单词的长度，因为每个单词都一样长
	wordLen := len(words[0])

	// 计算包含子串的窗口大小
	window := wordLen * len(words)

	// i只需要扫描过1个单词的每个位置即可，j根据窗口大小进行遍历就都会扫描到了
	for i := 0; i < wordLen; i++ {
		for j := i; j <= len(s) - window; j += wordLen {
			tmpStr := s[j : j+window]
			temMap := make(map[string]int)

			// 从最尾巴开始比对字符串
			// 将扫过的字符串放到map中，和目标的对比，如果出现大于的计数则认为不符合，j直接跳到不符合的那里再次开始
			// 如果遍历到了第一位都符合，则j符合条件
			for k := len(words) - 1; k >= 0; k-- {
				word := tmpStr[k*wordLen : (k+1)*wordLen]

				count := temMap[word] + 1
				if count > wordMap[word] {
					j = j + k*wordLen
					break
				} else if k == 0 {
					ret = append(ret, j)
				} else {
					temMap[word] = count
				}
			}
		}
	}

	return ret
}
