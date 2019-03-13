package main

import "fmt"

/**

题目:
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true

*/

func main() {
	str := "{[]} "
	r := isValid(str)
	fmt.Println(r)
}

/**
方法1:

这道题的关键点在于，左边是可以直接添加，但是右边添加的话，是需要左边已存在配对的左边

可以采用消除法进行，添加右边的时候，由于例子4所以，右边第一个肯定是其配对否则就需要报错

时间复杂度：O(n)
空间复杂度：O(1)
*/
func isValid(str string) bool {
	temp := map[byte]byte{
		'}':'{',
		')':'(',
		']':'[',
	}

	var list []byte
	for i := range str {
		curStr := str[i]
		switch curStr {
		case '(','{','[':
			list = append(list, curStr)
		case ')','}',']':
			if len(list) == 0 {
				return false
			}

			if list[len(list)-1] == temp[curStr] {
				list = list[:len(list)-1]
			} else {
				return false
			}
		}
	}


	return len(list) == 0
}
