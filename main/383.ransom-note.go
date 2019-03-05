package main

import "fmt"

/**

题目:
给定一个赎金信 (ransom) 字符串和一个杂志(magazine)字符串，
判断第一个字符串ransom能不能由第二个字符串magazines里面的字符构成。
如果可以构成，返回 true ；否则返回 false。

(题目说明：为了不暴露赎金信字迹，要从杂志上搜索各个需要的字母，组成单词来表达意思。)

注意：
你可以假设两个字符串均只含有小写字母。

canConstruct("a", "b") -> false
canConstruct("aa", "ab") -> false
canConstruct("aa", "aab") -> true

*/

func main() {
	ransom := "aa"
	magazine := "aab"
	r := canConstruct(ransom,magazine)
	fmt.Println(r)
}

/**
方法1:

因为只含有小写字母，所以提供两种方法，都需要用数组或者map：
1.来一个map[rune]int的map，这样好处是rune可以支持中文，遍历一遍杂志的，存在的字符码+1，后面遍历信的文字，进行减1操作，当小于0的时候则return false；
2.来一个[]int的数组，因为ascii只有256位，所以直接遍历一遍杂志的，存在的字符ascii+1，后面遍历信的文字，进行减1操作，当小于0

时间复杂度：O(n)
空间复杂度：O(1)
*/
func canConstruct(ransomNote string, magazine string) bool {
	temp := make([]int, 256)
	for _,v := range magazine {
		temp[v]++
	}

	fmt.Println(temp)
	for _,v := range ransomNote {
		temp[v]--
		if temp[v] < 0 {
			return false
		}
	}
	return true
}
