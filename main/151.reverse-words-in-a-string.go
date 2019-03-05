package main

import (
	"fmt"
	"strings"
)

/**

题目:
给定一个字符串，逐个翻转字符串中的每个单词。

示例 1：
输入: "the sky is blue"
输出: "blue is sky the"

示例 2：
输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

说明：
无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

进阶：
请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。

*/

func main() {
	str := " hello world "
	str = "a good   example"
	r := reverseWords(str)
	fmt.Println(r)
}

/**
方法1:

先去掉左右空格，再将三个或者两个空格转为一个，再整体翻转，再单个翻转

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseWords(s string) string {
	// 去除两边空字符串
	s = strings.Trim(s, " ")
	n:=-1
	s = strings.Replace(s, "   ", " ", n)
	s = strings.Replace(s, "  ", " ", n)


	// 整体翻转
	sLen := len(s)
	s = reverse(s,0,sLen-1)

	// 取出空字符串
	i := 0
	j := 0
	for j<sLen {
		for j < sLen && s[j] != ' ' {
			j++
		}

		s = reverse(s,i,j-1)
		j++
		i = j
	}

	return s
}

func reverse(s string, i int, j int) string {
	sArr := []byte(s)
	for i < j {
		sArr[i], sArr[j] = s[j], s[i]
		i++
		j--
	}
	return string(sArr)
}
