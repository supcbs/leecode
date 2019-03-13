package main

import "fmt"

/**

题目:
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:
输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"

示例 2:
输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

*/

func main() {
	str := "hello"
	r := temp1(str)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func longestValidParentheses(s string) int {
	return 1
}


func longestValidParenthesesOK(s string) int {
	left := 0
	lresult := 0
	llastlen := 0
	lmax := 0
	for i := 0; i < len(s); i++ {
		if s[i] ==  '(' {
			left++
			lresult++
		} else if s[i] ==  ')' && left > 0 {
			left--
			lresult++
			if left == 0 {
				if lresult + llastlen > lmax {
					lmax = lresult + llastlen
				}
				llastlen = lresult + llastlen
				lresult = 0
			}
		} else {
			llastlen = 0
			lresult = 0
		}
	}

	if left == 0 {
		if lresult + llastlen > lmax {
			lmax = lresult + llastlen
		}
	}

	right := 0
	rresult := 0
	rlastlen := 0
	rmax := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] ==  ')' {
			right++
			rresult++
		} else if s[i] ==  '(' && right > 0 {
			right--
			rresult++
			if right == 0 {
				if rresult + rlastlen > rmax {
					rmax = rresult + rlastlen
				}
				rlastlen = rresult + rlastlen
				rresult = 0
			}
		} else {
			rlastlen = 0
			rresult= 0
		}
	}

	if right == 0 {
		if rresult + rlastlen > rmax {
			rmax = rresult + rlastlen
		}
	}

	max := lmax
	if rmax > max {
		max = rmax
	}

	return max
}