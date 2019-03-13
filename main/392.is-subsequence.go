package main

import "fmt"

/**

题目:
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
你可以认为 s 和 t 中仅包含英文小写字母。
字符串 t 可能会很长（长度 ~= 500,000），而 s 是个短字符串（长度 <=100）。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。

示例 1:
s = "abc", t = "ahbgdc"
返回 true.

示例 2:
s = "axc", t = "ahbgdc"
返回 false.

后续挑战 :
如果有大量输入的 S，称作S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T 的子序列。在这种情况下，你会怎样改变代码？

*/

func main() {
	s := "abc"
	t := "ahbgdc"

	r := isSubsequence(s, t)
	fmt.Println(r)
}

/**
方法1:

这道题是必须要遍历t的。
遍历t的时候，进行s与t的比对，如果相等，则i++，进行下一位的比较
这里的阻断条件是i==s的长度的时候

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}

	i := 0
	for j := 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
			if i == len(s) {
				return true
			}
		}
	}
	return false
}
