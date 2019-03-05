package main

import "fmt"

/**

题目:
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:
s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.

*/

func main() {
	str := "leetcode"
	r := firstUniqChar(str)
	fmt.Println(r)
}

/**
方法1:

因为只有26个小写字母，所以可以直接存在一个arr里头，arr，
再遍历一遍字符串，若果等于1则认为不重复

---
golang可以直接int('字母')会转为ascii码

时间复杂度：O(n)
空间复杂度：O(26)
*/
func firstUniqChar(s string) int {
	temp := make([]int,26)
	for _,v := range s {
		temp[int(v-'a')]++
	}

	ret := -1
	for k,v := range s {
		if temp[int(v-'a')] == 1 {
			ret = k
			break
		}
	}

	return ret
}
