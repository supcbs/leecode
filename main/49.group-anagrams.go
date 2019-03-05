package main

import "fmt"

/**

题目:
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:
输入: ["eat", "tea", "tan", "ate", "nat", "bat"],

输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

说明：
所有输入均为小写字母。
不考虑答案输出的顺序。

*/

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	r := groupAnagrams(arr)
	fmt.Println(r)
}

/**
方法1:

思路就是将字符串按照英文字母进行排序，然后当作map的key，value由i递增

当遇到相同的key则直接加到相同的返回数组里头

时间复杂度：O(n)
空间复杂度：O(1)
*/
func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	if len(strs) == 0 {
		return ret
	}

	temp := make(map[string]int)

	i := 0
	for _, str := range strs {
		rStr := reorderStr(str)
		if k, ok := temp[rStr]; ok {
			ret[k] = append(ret[k], str)
		} else {
			temp[rStr] = i
			ret = append(ret, []string{str})
			i++
		}
	}

	return ret
}

func reorderStr(s string) string {
	var temp [26]int
	for i := 0; i < len(s); i++ {
		temp[s[i]-byte('a')]++
	}

	r := make([]byte, len(s))
	k := 0
	// 这里十分注意是每一位都递增，所以需要单独来一个k
	for i := 0; i < 26; i++ {
		for j := 0; j < temp[i]; j++ {
			r[k] = byte(i) + byte('a')
			k++
		}
	}
	fmt.Println(string(r))
	return string(r)
}
