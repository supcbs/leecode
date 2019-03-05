package main

import (
	"fmt"
)

/**

题目:
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

*/

func main() {
	str := []string{"b","b"}
	r := longestCommonPrefix(str)
	fmt.Println(r)
}

/**
方法1:

先找出前面的公共部分，在寻找后面的进行比对.
碰到任何一个不匹配则将pos置为j-1，若为-1，后续起始也不用匹配了


时间复杂度：O(n)
空间复杂度：O(1)
*/
func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen <= 0 {
		return ""
	}

	if strsLen <= 1 {
		return strs[0]
	}

	// 先取出最短的pos
	minPos := len(strs[0])
	for i:=0;i<strsLen;i++{
		if minPos > len(strs[i]) {
			minPos = len(strs[i])
		}
	}
fmt.Println(minPos)
	for i := 1; i < strsLen; i++ {
		var j int
		// 遍历的时候，只遍历短的为准

		for j = 0; j < minPos; j++ {
			fmt.Println(i,j)
			if strs[0][j] != strs[i][j] {
				minPos = j
			}
		}
	}

	if minPos == 0 || minPos == -1 {
		return ""
	}

	return strs[0][:minPos] //左开右闭，所以需要
}
