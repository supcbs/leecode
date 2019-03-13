package main

import "fmt"

/**

题目:

找到给定字符串（由小写字符组成）中的最长子串 T ，
要求 T 中的每一字符出现次数都不少于 k 。输出 T 的长度。

示例 1:
输入:
s = "aaabb", k = 3
输出:
3

最长子串为 "aaa" ，其中 'a' 重复了 3 次。

示例 2:
输入:
s = "ababbc", k = 2
输出:
5
最长子串为 "ababb" ，其中 'a' 重复了 2 次， 'b' 重复了 3 次。

*/

func main() {
	s := "bbaaa"
	k := 3
	r := longestSubstring(s,k)
	fmt.Println(r)
}

/**
方法1:

1.先进行字符串个数的累加，看一下最大的累加数和k比，是否小于
2.然后遍历字符每个字母，计算出满足条件的子字符串数组
3.然后遍历子字符串数组，进行递归调用，判断得出最大的值
4.
5.

时间复杂度：O(?)
空间复杂度：O(1)
*/
func longestSubstring(s string, k int) int {
	if k == 0 {
		return len(s)
	}

	m := make(map[rune]int)
	maxv := 0
	// 给每一个字母进行累加，同时算出累加数最大的字符
	for _,v := range s {
		m[v]++
		if m[v] > maxv {
			maxv = m[v]
		}
	}

	// 如果累加的最大数都没有k大，则认为没有符合条件的值
	if maxv < k {
		return 0
	}

	ss := make([]string,0)

	start := 0
	all := 1

	for i,v := range s {
		if m[v] < k {
			// 计算两个不符合条件的字符，中间间隔的长度是否大于k
			// 大于则认为这中间可能存在符合条件的最长子串
			if len(s[start:i]) >= k {
				ss = append(ss, s[start:i])
			}
			start = i +1
			all = 0
		}
	}

	ss = append(ss, s[start:])
	// 如果都是大于k的，则直接返回字符串长度的就好
	if all == 1 {
		return len(s)
	}

	maxv = 0
	for _,v := range ss {
		if len(v) < k {
			continue
		}

		t := longestSubstring(v, k)
		if t > maxv {
			maxv = t
		}
	}
	return maxv
}
