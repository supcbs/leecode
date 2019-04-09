package main

import "fmt"

/**

题目，实现php中strpos

从字符串中寻找目标字符串，找到则返回索引位置，没有找到则返回-1。

真正讲懂：为什么要向右移动4位呢，因为移动4位后，模式串中又有个“AB”可以继续跟S[8]S[9]对应着
https://blog.csdn.net/dl962454/article/details/79910744

https://blog.csdn.net/buppt/article/details/78531384
https://blog.csdn.net/starstar1992/article/details/54913261/
*/

func main() {
	str := "bacbababadababacambabacaddababacasdsd"
	t := "ababaca"
	r := kmp(str,t)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n+m)
空间复杂度：O(1)
*/
func kmp(str string, targetStr string) int {
	strLen := len(str)
	targetStrLen := len(targetStr)

	next := make([]int, targetStrLen)
	next = calNext(targetStr)

	k := -1
	for i := 0; i < strLen; i++ {
		for k > -1 && targetStr[k+1] != str[i] {
			k = next[k]
		}

		if targetStr[k+1] == str[i] {
			k++
		}

		if k == targetStrLen-1 {
			return i - targetStrLen + 1
		}
	}

	return -1
}

// str := "bacbababadababacambabacaddababacasdsd"
//	t := "ababaca"
/**
   a   b   a   b   a   c   a
   -1  -1  0   1   2   -1
 */
func calNext(str string) []int {
	next := make([]int, len(str))

	next[0] = -1
	k := -1

	for i := 1; i < len(str); i++ {
		for k > -1 && str[k+1] != str[i] {
			k = next[k]
		}

		if str[k+1] == str[i] {
			k++
		}
		next[i] = k
	}

	return next
}
