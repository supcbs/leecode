package main

import "fmt"

/**

题目:
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/

func main() {
	n := 3
	r := generateParenthesis(n)
	fmt.Println(r)
}

/**
方法1:

这道题使用回溯法

关键点在于：
1.左边和右边剩余的数量都为0的时候则这个字符串可以加到返回里头
2.当左边>0那么，将字符串加上'('，继续进行递归
3.当右边>0 & 大于左边的时候，将字符串加上)然后继续递归

时间复杂度：O(n)
空间复杂度：O(1)
*/
func generateParenthesis(n int) []string {
	var ret []string
	g(&ret, "",n,n)
	return ret
}

func g(ret *[]string,s string, left int,right int) {
	if left == 0 && right == 0 {
		*ret = append(*ret,s)
		return
	}

	fmt.Println(s)
	if left > 0 {
		g(ret, s+"(", left-1,right)
	}

	if right > 0 && right > left {
		g(ret, s+")", left,right-1)
	}
}
