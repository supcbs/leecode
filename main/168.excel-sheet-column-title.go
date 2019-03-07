package main

import "fmt"

/**

题目:
给定一个正整数，返回它在 Excel 表中相对应的列名称。
例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...

示例 1:
输入: 1
输出: "A"

示例 2:
输入: 28
输出: "AB"

示例 3:
输入: 701
输出: "ZY"

*/

func main() {
	n := 1
	r := convertToTitleOK(n)
	fmt.Println(r)
}

/**
方法1:



时间复杂度：O()
空间复杂度：O(1)
*/
func convertToTitle(n int) string {
	if n <= 0 {
		return ""
	}
	res := make([]byte, 0)

	for n != 0 {
		n--
		c:=n%26
		res = append(res, byte(c)+'A')
		n /= 26
	}

	ret := make([]byte, 0)
	for i:=len(res)-1;i>=0;i-- {
		ret = append(ret, res[i])
	}
	return string(ret)
}

func convertToTitleOK(n int) string {
	if n <= 0 {
		return ""
	}

	ans := ""
	for n > 0 {
		n --
		c := byte(n%26) + 'A'
		ans = fmt.Sprintf("%c%s", c, ans)
		n /=26
	}

	return ans
}
