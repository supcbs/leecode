package main

import "fmt"

/**

题目:
给定不同面额的硬币 coins 和一个总金额 amount。
编写一个函数来计算可以凑成总金额所需的最少的硬币个数。
如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1

示例 2:
输入: coins = [2], amount = 3
输出: -1

*/

func main() {
	coins := []int{1,2,5}
	amount := 11
	r := coinChange(coins, amount)
	fmt.Println(r)
}

/**
方法1:

这道题已经不能用贪心算法了。
所以只能用动态规划。

dp[i] = min(dp[i], dp[amount-coins[j]] + 1)
i代表钱数。其值代表硬币个数。
硬币只有使用和不使用之分，dp[i]代表未使用当前币，dp[amount-coins[j]] + 1代表使用了

时间复杂度：O(n)
空间复杂度：O(1)
*/
func coinChange(coins []int, amount int) int {
	if len(coins) <= 0 {
		return -1
	}

	if amount <= 0  {
		return 0
	}


	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _,coin := range coins {
		for i := coin;i<=amount;i++{
			if dp[i] > dp[i-coin] + 1 {
				dp[i] = dp[i-coin] + 1
			}
		}
	}

	fmt.Println(dp)

	if dp[amount] > amount {
		return -1
	} else {
		return dp[amount]
	}
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
