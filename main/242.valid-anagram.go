package main

import "fmt"

/**

题目:
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true

示例 2:
输入: s = "rat", t = "car"
输出: false

说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？

*/

func main() {
	str := "rat"
	t := "car"
	r := isAnagram(str,t)
	fmt.Println(r)
}

/**
方法1:

来一个map，第一个字符串先存一把，，遍历第二个字符串进行--，发现小于0则不行

时间复杂度：O(n+m)
空间复杂度：O(1)
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	temp := make(map[byte]int)
	for i:=0 ; i <= len(s) - 1; i++{
		temp[s[i]]++
		temp[t[i]]--
	}
	for _,v := range temp {
		if v != 0 {
			return false
		}
	}

	return true
}


func isAnagramOK(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	b := make([]int, 26)
	for _, r := range s {
		b[r-'a']++
	}
	for _, r := range t {
		b[r-'a']--
		if b[r-'a'] < 0 {
			return false
		}
	}
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}