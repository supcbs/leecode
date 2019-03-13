package main

import "fmt"

/**

题目:
给定一个字符串 S 和一个字符串 T，请在 S 中找出包含 T 所有字母的最小子串。

示例：
输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"

说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

*/

func main() {
	//r := minWindow("ADOBECODEBANC", "ABC")
	r := minWindow("a", "a")
	fmt.Println(r)
}

/**
方法1:

和209类似,3题
都是采用滑动窗口+双指针的方式进行
1.满足窗口的值 ： 即正好都包含了目标字符，可以通过计数等于目标字符长度的方式来判断，或者遍历
2.ij指针往前走的时机，每次j都往前

关键点在于i合适往前，当即使失去当前数也能满足条件的时候就可以往前，
此时为了节省时间需要一个map记录下来该字母出现的个数，与要求的字母map进行比较必须要>才行

3.最后每次循环都更新最短串就就好了

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)

	// ascii就只有128位
	needMap := [128]int{}
	haveMap := [128]int{}

	for _, v := range t {
		needMap[v]++
	}

	ret := ""
	haveCount := 0
	minLen := sLen
	for i, j := 0, 0; j < sLen; j++ {
		// 判断拥有数
		if haveMap[s[j]] < needMap[s[j]] {
			haveCount++
		}
		haveMap[s[j]]++

		// i向前计划,保证i尽可能短
		for i < j && haveMap[s[i]] > needMap[s[i]] {
			haveMap[s[i]]--
			i++
		}

		tempLen := j - i + 1
		if haveCount == tLen && tempLen <= minLen {
			minLen = tempLen
			ret = s[i : j+1]
		}

	}

	return ret
}
//"ADOBECODEBANC", "ABC"