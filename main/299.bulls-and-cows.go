package main

import (
	"fmt"
	"strconv"
)

/**

题目:
你正在和你的朋友玩 猜数字（Bulls and Cows）游戏：你写下一个数字让你的朋友猜。
每次他猜测后，你给他一个提示，告诉他有多少位数字和确切位置都猜对了（称为“Bulls”, 公牛），
有多少位数字猜对了但是位置不对（称为“Cows”, 奶牛）。你的朋友将会根据提示继续猜，直到猜出秘密数字。
请写出一个根据秘密数字和朋友的猜测数返回提示的函数，用 A 表示公牛，用 B 表示奶牛。

请注意秘密数字和朋友的猜测数都可能含有重复数字。

示例 1:
输入: secret = "1807", guess = "7810"
输出: "1A3B"
解释: 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。

示例 2:
输入: secret = "1123", guess = "0111"
输出: "1A1B"
解释: 朋友猜测数中的第一个 1 是公牛，第二个或第三个 1 可被视为奶牛。

说明: 你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。
*/

func main() {
	secret := "1807"
	guess := "7810"
	r := getHint(secret,guess)
	fmt.Println(r)
}

/**
方法1:
很容易想到的是必须要遍历一次。
判断公牛则相对简单，直接每一位都遍历，相等即为公牛。
判断奶牛需要借助一个额外的map，由于单个数字的取值范围为0-9，标记0-9两个字符串出现的次数，通过数量比对来得出奶牛的数。

时间复杂度：O(n)
空间复杂度：O(1)
---
这道题用Golang实现的时候还考了一个细节：
Golang的字符串是UTF8的unicode文本，也就是string[i]出来的不是字符串而是字符串对应的ASCII码。
所以需要格式化城字符串，在将字符串转为int类型
*/
func getHint(secret string, guess string) string {
	var A,B,i,j int
	AMap := make([]int, 10)
	BMap := make([]int, 10)

	secretLen := len(secret)
	for i < secretLen {
		secretS := fmt.Sprintf("%c", secret[i])
		guessS := fmt.Sprintf("%c", guess[i])

		secretInt,_ := strconv.Atoi(secretS)
		guessInt,_ := strconv.Atoi(guessS)

		if secretInt == guessInt {
			A ++
		} else {
			AMap[secretInt] ++
			BMap[guessInt] ++
		}
		i++
	}

	for j<10 {
		if AMap[j] < BMap[j] {
			B += AMap[j]
		} else {
			B += BMap[j]
		}
		j++
	}

	return fmt.Sprint(A,"A",B,"B")
}

