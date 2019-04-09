package main

import "fmt"

/**

题目:
给定三个字符串 s1, s2, s3, 验证 s3 是否是由 s1 和 s2 交错组成的。

示例 1:
输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
输出: true

示例 2:
输入: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
输出: false

*/

func main() {
	r := isInterleave("aabcc","dbbca","aadbbcbcac")
	fmt.Println(r)
}

/**
方法1:
[-1,-1,0,1,2,-1,0]

bacbababadababacambabacaddababacasdsd

ababaca

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	return false
}
























