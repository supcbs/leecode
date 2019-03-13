package main

import "fmt"

/**

题目:

给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。
找到并返回可以用这种方式转换的最短回文串。

示例 1:
输入: "aacecaaa"
输出: "aaacecaaa"

示例 2:
输入: "abcd"
输出: "dcbabcd"

*/

func main() {
	str := "abcd"
	r := shortestPalindrome(str)
	fmt.Println(r)
}

/**
方法1:

和第5题对应，起始也有点不太像
---
这道题题目的意思是说只能往前头加字符串，也就是说，
前头需要加的字符串就是非回文的部分。

1.首先从尾巴开始比对，ij指针，如果都相同则需要复制的尾巴位数为0，如果有不一样，则从尾巴前一位开始比对，计算出这个需要加到位数
2.从尾巴开始遍历位数，加到一个字符串，注意是倒序
3.第二步得到的字符串添加到原字符串开头

时间复杂度：O(n)
空间复杂度：O(1)
*/
func shortestPalindrome(s string) string {
	if len(s) <= 0 {
		return s
	}

	count := 0
	for ; count < len(s); count++ {
		ok := true

		// 双指针滚动
		j := 0
		for i := len(s) - 1 - count; i >= 0; i-- {
			if s[i] != s[j] {
				ok = false
				break
			}
			j++
		}

		// 代表已经是回文了
		if ok {
			break
		}
	}

	res := []byte("")
	for i := len(s) - 1; i > len(s)-1-count; i-- {
		res = append(res, s[i])
	}

	ret := append(res, []byte(s)...)
	return string(ret)
}
