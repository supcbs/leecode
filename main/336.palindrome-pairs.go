package main

import "fmt"

/**

题目:
给定一组唯一的单词， 找出所有不同 的索引对(i, j)，使得列表中的两个单词，
words[i] + words[j] ，可拼接成回文串。

示例 1:
输入: ["abcd","dcba","lls","s","sssll"]
输出: [[0,1],[1,0],[3,2],[2,4]]
解释: 可拼接成的回文串为 ["dcbaabcd","abcddcba","slls","llssssll"]

示例 2:
输入: ["bat","tab","cat"]
输出: [[0,1],[1,0]]
解释: 可拼接成的回文串为 ["battab","tabbat"]

*/

func main() {
	str := "hello"
	r := temp1(str)
	fmt.Println(r)
}

/**
方法1:

循环遍历每个字符串拼装之后进行回文判断。

时间复杂度：O(n)
空间复杂度：O(1)
*/
func palindromePairs(words []string) [][]int {

}
