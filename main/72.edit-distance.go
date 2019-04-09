package main

import (
	"fmt"
)

/**

题目:
给定两个单词 word1 和 word2，计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符

示例 1:
输入: word1 = "horse", word2 = "ros"
输出: 3
解释:
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')

示例 2:
输入: word1 = "intention", word2 = "execution"
输出: 5
解释:
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')

*/

func main() {
	r := minDistance("hello", "hllo")
	fmt.Println(r)
}

/**
方法1:
	//建立数组dp[][]来存储 以word1[i]为结尾的字符串 转换成 以word2[j]为结尾的字符串 所需的最小操作数
    //当新遍历一对来自word1，word2的字符
    //若 word1[i] == word2[j] 表示不需要操作，则dp[i][j]=dp[i-1][j-1]
    //若 word1[i] != word2[j] 则可以有三种情况
    //   1、替换 word1[i] 把 word1[i] 替换成 word2[j] 需要 dp[i-1][j-1]+1步
    //   2、删除 word1[i] 把 word1[i] 删除成 word1[i-1] 需要 dp[i][j-1]+1步
    //   3、删除 word2[j] 把 word2[j] 删除成 word2[j-1] 需要 dp[i-1][j]+1步(增加word1和删除word2一个效果)
    // 取这三个中最小值

有个位置特别关键
前一个位置     | i剪掉一位的位置
------------------------------------
j减掉一位的位置 | 当前的位置（增加了一位）

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minDistance(word1 string, word2 string) int {
	len1 := len(word1)
	len2 := len(word2)

	var dp [][]int
	for i := 0; i < len1+1; i++ {
		tmp := make([]int, len2+1)
		dp = append(dp, tmp)
	}

	for i := 1; i <= len1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len2; i++ {
		dp[0][i] = i
	}
/**
[0 1 2 3 4]
[1 0 1 2 3]
[2 1 1 2 3]
[3 2 1 1 2]
[4 3 2 1 2]
[5 4 3 2 1]
 */
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
		}
	}

	for _,v:=range dp{
		fmt.Println(v)
	}

	return dp[len1][len2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
