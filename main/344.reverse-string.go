package main

import "fmt"

/**

题目:
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。

示例 1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]

示例 2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/

func main() {
	str := "hello"
	reverseString([]byte(str))
	fmt.Println(str)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func reverseString(s []byte) string {
	sLen := len(s)

	l := 0
	r := sLen - 1
	for l<r {
		s[l],s[r] = s[r],s[l]
		fmt.Println(s[l],s[r])
		l++
		r--
	}

	return string(s)
}


func reverseString2(s []byte)  {
	// 边界条件
	if len(s) <= 0 {
		return
	}

	// 初始值设置
	i := 0
	j := len(s) - 1

	recursion(s, i, j)

	// // 交换逻辑
	// for i <= j {
	//     s[i], s[j] = s[j], s[i]
	//     i++
	//     j--
	// }

	return
}

func recursion(s []byte, i, j int)  {
	if i >= j {
		return
	}
	s[i], s[j] = s[j], s[i]
	i++
	j--

	recursion(s, i, j)
	return
}
