package main

import (
	"fmt"
	"strings"
)

/**

题目:
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，
在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。
如果不存在，则返回  -1。

示例 1:
输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:
输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

*/

func main() {
	haystack := "he你ll我o"
	needle := "我o"
	r := strStrOK(haystack, needle)
	fmt.Println(r)
}

/**
方法1:
思路很简单，遍历haystack每一个字母，如果等于needle第一个字母则继续，否则继续遍历haystack下一个字母
---
需要注意的是:如果包含汉字，直接用for range比for ；；省事，前者是unicode后者是ascii码

时间复杂度：O(n)
空间复杂度：O(1)
*/
func strStr(haystack string, needle string) int {
	if len(needle) <= 0{
		return 0
	}

	for i:=0;i<len(haystack);i++ {
		flag := true

		start := i
		if len(haystack) - i <= len(needle){
			break
		}
		for j:=0;j<len(needle);j++ {
			if haystack[start] != needle[j] {
				flag = false
				break
			}
			start++
		}
		if flag {
			return i
			break
		}
	}

	return -1
}

/**
方法2:

官方题解：使用了strings包
这里需要注意的是默认使用的是ascii需要转成unicode码的序号


时间复杂度：O(n)
空间复杂度：O(1)
*/
func strStrOK(haystack string, needle string) int {
	result := strings.Index(haystack, needle)
	if result >= 0{
		prefix := []byte(haystack[0:result])
		res := []rune(string(prefix))
		result = len(res)
	}
	return result
}