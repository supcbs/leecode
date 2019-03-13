package main

import "fmt"

/**

题目:
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:
输入: 121
输出: true

示例 2:
输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

示例 3:
输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。

*/

func main() {
	num := -121
	r := isPalindrome(num)
	fmt.Println(r)
}

/**
方法1:

这个高级的做法是不变成字符串，直接用数字进行计算。

每次将最低位去除，同时复值给y，y需要已累计到的数*10之后加上x的最低位
进行比较。当x=y,或者说x/10=y也行。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}

	y := 0
	r := 0
	for y < x {
		r = x % 10
		x = x / 10
		y = y*10 + r

		if x == y || x/10 == y {
			return true
		}
	}

	return false
}
