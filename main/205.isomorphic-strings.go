package main

import "fmt"

/**

题目:
给定两个字符串 s 和 t，判断它们是否是同构的。
如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。

示例 1:
输入: s = "egg", t = "add"
输出: true

示例 2:
输入: s = "foo", t = "bar"
输出: false

示例 3:
输入: s = "paper", t = "title"
输出: true
说明:
你可以假设 s 和 t 具有相同的长度。

*/

func main() {
	str1 := "foo"
	str2 := "bar"
	r := isIsomorphic(str1, str2)
	fmt.Println(r)
}

/**
方法1:

这道题定义两个map,一个存储a->b的映射，另一个存储b->a的映射
1.当两个map都不存在的时候，直接存储映射
2.存在情况不同，直接返回false
3.都存在情况下，比较value只是否与当前的相同，相同则跳过，不同直接返回false


时间复杂度：O(n)
空间复杂度：O(1)
*/
func isIsomorphic(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)

	if sLen != tLen {
		return false
	}

	if sLen == 0 {
		return true
	}

	aToB := make(map[byte]byte)
	bToA := make(map[byte]byte)

	for i:= 0 ; i < sLen;i++{
		v1,ok1 := aToB[s[i]]
		v2,ok2 := bToA[t[i]]

		if ok1 != ok2 {
			return false
		}

		if ok1 {
			if v1 == t[i] && v2 == s[i] {
				continue
			}
			return false
		}

		aToB[s[i]],bToA[t[i]] = t[i],s[i]
	}

	return true
}







