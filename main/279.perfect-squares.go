package main

import (
	"fmt"
	"math"
)

/**

题目:
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。
你需要让组成和的完全平方数的个数最少。

示例 1:
输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.

示例 2:
输入: n = 13
输出: 2
解释: 13 = 4 + 9.

*/

func main() {
	n := 13
	r := numSquares(n)
	fmt.Println(r)
}

/**
方法1:

dp(i) = min(dp(i),dp(num-j*j)+1)

i的范围是从1开始，到n，所以数组需要n+1的长度

时间复杂度：O(n)
空间复杂度：O(1)
*/
func numSquares(n int) int {
	if n <= 0 {
		return 1
	}
	numMax := int(^uint(0) >> 1)

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = numMax
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			if dp[i] > dp[i-j*j]+1 {
				dp[i] = dp[i-j*j] + 1
			} else {
				dp[i] = dp[i]
			}
		}
	}
	return dp[n]
}

/**
本质上是四平方和的原理。

但是还是走动态规划的方式，适合走上面那一步
 */
func numSquaresOK(n int) int {
	for n%4 == 0 {
		n/=4
	}

	if n%8 == 7 {
		return 4
	}
	a:=0
	for a*a <= n {
		b := int(math.Sqrt(float64(n-a*a)))
		if a*a+b*b ==n {
			if a!=0 && b!=0 {
				return 2
			} else if a==0 && b==0 {
				return 0
			} else {
				return 1
			}
		}
		a++
	}

	return 3
}