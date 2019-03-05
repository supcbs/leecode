package main

import "fmt"

/**

题目:
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:
输入: "Hello World"
输出: 5

*/

func main() {
	str := "hello world"
	r := lengthOfLastWord(str)
	fmt.Println(r)
}

/**
方法1:

思路很简单，直接从尾巴开始找，但是需要注意的是需要去掉尾巴的空字符串

时间复杂度：O(n)
空间复杂度：O(1)
*/
func lengthOfLastWord(s string) int {
	sLen := len(s)
	if sLen <= 0 {
		return 0
	}

	idx := sLen - 1
	// 去掉空格
	for ; idx >= 0 && s[idx] == ' '; idx-- {

	}
fmt.Println(idx)
	// 数出最后一个字符串的长度
	count := 0
	for i := idx; i >= 0 && s[i] != ' '; i-- {
		count++
	}

	return count
}
