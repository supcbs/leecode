package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**

题目:

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:
输入: "A man, a plan, a canal: Panama"
输出: true

示例 2:
输入: "race a car"
输出: false

*/

func main() {
	str := "\"0P\""
	r := isPalindromeOK(str)
	fmt.Println(r)
}

/**
方法1:

理解题意很重要，
回文字符串指的是，整个字符串是轴对称的，
也就是说两个指针从头和尾巴向中间跑的时候，移动相同的位置则字符应该是一样的

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) {
			i++
			continue
		}
		if !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) {
			j--
			continue
		}
		fmt.Println(i, j)
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

/**
0 29
2 28
3 27
4 26
7 25
9 24
10 21
11 20
12 19
15 18
*/
func isPalindromeOK(s string) bool {
	if len(s) <= 0 {
		return true
	}

	s = strings.ToLower(s)
	fmt.Println(s, len(s))
	for i, j := 0, len(s)-1; i < j; {
		if !(s[i] <= 'z' && s[i] >= 'a') && !(s[i] <= '9' && s[i] >= '0') {
			i++
			continue
		}

		if !(s[j] <= 'z' && s[j] >= 'a') && !(s[j] <= '9' && s[j] >= '0') {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
