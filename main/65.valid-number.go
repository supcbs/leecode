package main

import (
	"fmt"
	"strings"
)

/**

题目:
验证给定的字符串是否为数字。

例如:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true

Some examples:
"0" => true
" 0.1 " => true
"abc" => false
"1 a" => false
"2e10" => true
" -90e3   " => true
" 1e" => false
"e3" => false
" 6e-1" => true
" 99e2.5 " => false
"53.5e93" => true
" --6 " => false
"-+3" => false
"95a54e53" => false

说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。

更新于 2015-02-10:
C++函数的形式已经更新了。
如果你仍然看见你的函数接收 const char * 类型的参数，请点击重载按钮重置你的代码。

*/

func main() {
	str := "53.5e93"
	r := isNumber(str)
	fmt.Println(r)
}

/**
方法1:

规则简化后需要遵守几点：
1.e只能出现1次
2.符号值+-只能初心在第一个位置，或者前一位是e
3.e后面一定要有数字
4.点只能出现一次 && 不能在e之后
5.数字必须要出现

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isNumber(s string) bool {
	// 去除左右的空格
	s = strings.TrimSpace(s)
	if len(s) <= 0 {
		return false
	}

	numberSeen := false
	pointSeen := false
	eSeen := false
	numAfterE := true

	for k,v := range s {

		switch {
		case v >= '0' && v <= '9' :
			numberSeen = true
			numAfterE = true
		case v == '.':
			if eSeen || pointSeen {
				return false
			}
			pointSeen = true
		case v == 'e':
			if eSeen || !numberSeen {
				return false
			}
			eSeen = true
			numAfterE = false
		case v == '+' || v == '-':
			if k != 0 && s[k-1] != 'e' {
				return false
			}
			numberSeen = false
		default :
			return false
		}
	}

	return numberSeen && numAfterE
}
