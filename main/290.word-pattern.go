package main

import (
	"fmt"
	"strings"
)

/**

题目:
给定一种 pattern(模式) 和一个字符串 str ，判断 str 是否遵循相同的模式。
这里的遵循指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应模式。

示例1:
输入: pattern = "abba", str = "dog cat cat dog"
输出: true

示例 2:
输入:pattern = "abba", str = "dog cat cat fish"
输出: false

示例 3:
输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false

示例 4:
输入: pattern = "abba", str = "dog dog dog dog"
输出: false

说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。

*/

func main() {
	pattern := "abba"
	str := "dog dog dog dog"
	r := wordPattern(pattern,str)
	fmt.Println(r)
}

/**
方法1:

同205题，需要双向map

时间复杂度：O(n)
空间复杂度：O(1)
*/
func wordPattern(pattern string, str string) bool {
	if len(pattern) == 0{
		return false
	}

	if len(str) == 0 {
		return false
	}

	strArr := strings.Split(str, " ")
	if len(strArr) != len(pattern) {
		return false
	}

	pToS := make(map[byte]string)
	sToP := make(map[string]byte)
	for i:=0; i< len(strArr);i++{
		v1,ok1 := pToS[pattern[i]]
		v2,ok2 := sToP[strArr[i]]

		if ok1 != ok2 {
			return false
		}

		if ok1 {
			if v1 == strArr[i] &&  v2 == pattern[i] {
				continue
			}
			return false
		}

		sToP[strArr[i]], pToS[pattern[i]] = pattern[i],strArr[i]
	}

	return true
}
