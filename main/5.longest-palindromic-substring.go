package main

import "fmt"

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000

示例 1：
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。

示例 2：
输入: "cbbd"
输出: "bb"
*/

func main() {
	s := "babad"
	s = "cbbd"
	fmt.Println(longestPalindrome(s))
}

/**
方法1:

与214对应

时间复杂度：O(n)
空间复杂度：O(1)
*/
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin := 0
	maxLen := 1

	for i := 0; i < len(s); {
		// 剩余段比最大的长度的一半还小就不用比了
		if len(s)-i <= maxLen/2 {
			break
		}

		b, e := i, i

		// 若存在重复的则进行扩张
		for e < len(s)-1 && s[e+1] == s[e] {
			e++
		}
		i = e + 1

		// 中心往两边比对
		for e < len(s)-1 && b > 0 && s[b-1] == s[e+1] {
			b--
			e++
		}

		newLen := e - b + 1
		if newLen > maxLen {
			maxLen = newLen
			begin = b
		}
	}

	return s[begin : begin+maxLen]
}

/**
方法2:

这种解法先给每个字母的前后以及中间加上#，目的是为了让字符串都变成奇数
然后遍历一个中心，ij分别向两边扩张，算出最大的半径，即为最大的回文

时间复杂度：O(?)
空间复杂度：O(1)
*/
func longestPalindrome1(s string) string {
	slice := make([]string, 0, 4)

	for _, char := range s {
		slice = append(slice, "#")
		slice = append(slice, string(char))
	}

	slice = append(slice, "#") //先把字符串加上“#”

	maxR := 0     //记录最长的半径
	maxIndex := 0 //记录最长的 index
	sliceLen := len(slice)

	for index := range slice {

		if index >= 1 {

			r := 0
			i := index - 1
			j := index + 1

			for { //每一个字符 计算最长的半径

				if i < 0 || j >= sliceLen {
					break
				}

				if slice[i] == slice[j] {
					r++
					i--
					j++
				} else {
					break
				}

			}

			if r > maxR {
				maxR = r
				maxIndex = index
			}

		}
	}

	res := ""

	for index, str := range slice { //maxIndex-maxR到maxIndex+maxR是最长串

		if index >= (maxIndex-maxR) && index <= (maxIndex+maxR) && str != "#" {
			res += str
		}
	}

	return res
}
