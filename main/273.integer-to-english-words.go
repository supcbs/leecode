package main

import (
	"fmt"
	"strings"
)

/**

题目:
将非负整数转换为其对应的英文表示。可以保证给定输入小于 231 - 1 。

示例 1:
输入: 123
输出: "One Hundred Twenty Three"

示例 2:
输入: 12345
输出: "Twelve Thousand Three Hundred Forty Five"

示例 3:
输入: 1234567
输出: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"

示例 4:
输入: 1234567891
输出: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

*/

func main() {
	num := 123
	r := numberToWords(num)
	fmt.Println(r)
}

/**
方法1:

将所有的数字定义出来包括：10^n 和 20以下的数 和 几十位的

依次遍历，从10亿位开始依次向下。

时间复杂度：O(n)
空间复杂度：O(1)
*/
const (
	Billion  = 1000000000
	Million  = 1000000
	Thousand = 1000
	Hundred  = 100
)

var wordsDict = map[int]string{
	1:          "One",
	2:          "Two",
	3:          "Three",
	4:          "Four",
	5:          "Five",
	6:          "Six",
	7:          "Seven",
	8:          "Eight",
	9:          "Nine",
	10:         "Ten",
	11:         "Eleven",
	12:         "Twelve",
	13:         "Thirteen",
	14:         "Fourteen",
	15:         "Fifteen",
	16:         "Sixteen",
	17:         "Seventeen",
	18:         "Eighteen",
	19:         "Nineteen",
	20:         "Twenty",
	30:         "Thirty",
	40:         "Forty",
	50:         "Fifty",
	60:         "Sixty",
	70:         "Seventy",
	80:         "Eighty",
	90:         "Ninety",
	100:        "Hundred",
	1000:       "Thousand",
	1000000:    "Million",
	1000000000: "Billion",
}

func numberLT1000(num int) []string {
	var result []string
	d := num / 100
	if d > 0 {
		result = append(result, wordsDict[d], wordsDict[Hundred])
	}

	num %= 100
	if num == 0 {
		return result
	}

	if num < 20 {
		result = append(result, wordsDict[num])
	} else {
		t := num / 10
		a := num % 10
		result = append(result, wordsDict[t*10])
		if a > 0 {
			result = append(result, wordsDict[a])
		}
	}

	return result
}

func numberThousand(num int) []string {
	num /= Thousand
	if num > 0 {
		return append(numberLT1000(num), wordsDict[Thousand])
	}
	return nil
}

func numberMillion(num int) []string {
	num /= Million
	if num > 0 {
		return append(numberLT1000(num), wordsDict[Million])
	}
	return nil
}

func numberBillion(num int) []string {
	num /= Billion
	if num > 0 {
		return append(numberLT1000(num), wordsDict[Billion])
	}
	return nil
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}

	var ret []string
	if num >= Billion {
		ret = append(ret, numberBillion(num)...)
		num %= Billion
	}

	if num >= Million {
		ret = append(ret, numberMillion(num)...)
		num %= Million
	}

	if num >= Thousand {
		ret = append(ret, numberThousand(num)...)
		num %= Thousand
	}

	ret = append(ret, numberLT1000(num)...)

	return strings.Join(ret, " ")
}
