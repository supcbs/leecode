package main

import "fmt"

/**

题目:

编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:
输入: "hello"
输出: "holle"

示例 2:
输入: "leetcode"
输出: "leotcede"
说明:

元音字母不包含字母"y"。

*/

func main() {
	str := "hello"
	r := reverseVowels(str)
	fmt.Println(r)
}

/**
方法1:

双指针法：
left从左往右移动，碰到元音字母则停住
right从右往左移动，碰到元音字母则停住
left与right交换

临界条件：left <= right

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseVowels(s string) string {
	sb := []byte(s)
	l := 0
	r := len(sb) - 1
	for l < r {
		for l < r && !isV(sb[l]) {
			l ++
		}

		for l < r && !isV(sb[r]) {
			r --
		}
		sb[l],sb[r] = sb[r],sb[l]
		l++
		r--
	}

	return string(sb)
}

func isV(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		b = b - 'A' + 'a'
	}

	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}