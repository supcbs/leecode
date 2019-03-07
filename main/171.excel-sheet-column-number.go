package main

import (
	"fmt"
	"math"
)

/**

题目:
给定一个Excel表格中的列名称，返回其相应的列序号。

例如，

    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...

示例 1:
输入: "A"
输出: 1

示例 2:
输入: "AB"
输出: 28

示例 3:
输入: "ZY"
输出: 701

*/

func main() {
	str := "ZY"
	r := titleToNumber(str)
	fmt.Println(r)
}

/**
方法1:

思路很简单，就是遍历位数n*26^(len-i-1)

时间复杂度：O(n)
空间复杂度：O(1)
*/
func titleToNumber(s string) int {
	if len(s) == 0 {
		return 0
	}

	ret := 0
	for i := 0; i < len(s); i++ {
		initNum := len(s) - 1 - i
		tempNum := math.Pow(26, float64(initNum))
		ret += int(s[i]-'A'+1) * int(tempNum)
	}

	return ret
}
