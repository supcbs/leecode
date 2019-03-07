package main

import (
	"fmt"
	"math"
)

/**

题目:
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
例如， 罗马数字 2 写做 II ，即为两个并列的 1。
12 写做 XII ，即为 X + II 。
27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。
但也存在特例，例如 4 不写做 IIII，而是 IV。
数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。
同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

示例 1:
输入: 3
输出: "III"

示例 2:
输入: 4
输出: "IV"

示例 3:
输入: 9
输出: "IX"

示例 4:
输入: 58
输出: "LVIII"
解释: L = 50, V = 5, III = 3.

示例 5:
输入: 1994
输出: "MCMXCIV"
解释: M = 1000, CM = 900, XC = 90, IV = 4.

*/

func main() {
	str := 58
	r := intToRomanOK(str)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func intToRoman(num int) string {

	var luoNum = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	step := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	ret := ""
	for i := len(step) - 1; i >= 0; i-- {
		for num-step[i] >= 0 {
			ret += luoNum[step[i]]
			num -= step[i]
		}
	}

	return ret
}

/**
方法2:

官方最快解答
巧妙利用
2*n=1\10\100\1000
2*n+1=5\50\500
[2*n,2*n+1] = 4/40/400
[2*n,2*n+2] = 9/90/900

时间复杂度：O(n)
空间复杂度：O(1)
*/
func intToRomanOK(num int) string {
	trans := []byte{'I', 'V', 'X', 'L', 'C', 'D', 'M'}

	// 因为最大为1000，就是10的三次方
	x := 3
	ret := make([]byte, 0)
	for x >= 0 {
		rest := num / int(math.Pow10(x)) % 10
		if rest == 9 {
			ret = append(ret, trans[2*x], trans[2*x+2])
		} else if rest == 4 {
			ret = append(ret, trans[2*x], trans[2*x+1])
		} else if rest == 5 {
			ret = append(ret, trans[2*x+1])
		} else {
			for rest > 5 {
				ret = append(ret, trans[2*x+1])
				rest -= 5
			}

			for rest > 0 {
				ret = append(ret, trans[2*x])
				rest--
			}
		}

		x--
	}

	return string(ret)
}
