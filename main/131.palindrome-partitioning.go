package main

import "fmt"

/**

题目:
给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
返回 s 所有可能的分割方案。

示例:
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]

*/

func main() {
	str := "aab"
	r := partition(str)
	fmt.Println(r)
}

/**
方法1:

https://blog.csdn.net/weixin_34405557/article/details/87558869
深度优先

时间复杂度：O(n)
空间复杂度：O(1)
*/
func partition(s string) [][]string {
	var res [][]string
	var collect []string
	dfs(s, &collect, &res)
	return res
}

func dfs(s string, collect *[]string, res *[][]string) {
	if s == "" {
		collected := make([]string, len(*collect))
		copy(collected, *collect)
		*res = append(*res, collected)
		return
	}

	for i := 1; i < len(s); i++ {
		child := s[0:i]
		if isPalindromeString(child) {
			*collect = append(*collect, child)
			dfs(s[i:], collect, res)
			*collect = (*collect)[0 : len(*collect)-1]
		}
	}
}

func isPalindromeString(s string) bool {
	len := len(s)
	for i := 0; i < len/2; i++ {
		//字串可以直接[]，不需要避免charAt()
		if s[i] != s[len-1-i] {
			return false
		}
	}
	return true
}

func partitionOK(s string) [][]string {
	_isPalindrome := func(s string) bool {
		i, j := 0, len(s)-1
		for i < j {
			if s[i] != s[j] {
				return false
			}

			i, j = i+1, j-1
		}

		return true
	}

	ret := make([][]string, 0)
	arr := make([]string, 0)
	_walk := func(s string, fn interface{}) {
		if 0 == len(s) {
			cp := make([]string, len(arr))
			copy(cp, arr)
			ret = append(ret, cp)
			return
		}

		_fn := fn.(func(string, interface{}))
		for i := 0; i < len(s); i++ {
			str := s[0 : i+1]
			if _isPalindrome(str) {
				arr = append(arr, str)
				_fn(s[i+1:len(s)], fn)
				arr = arr[0 : len(arr)-1]
			}
		}
	}
	_walk(s, _walk)
	return ret
}
