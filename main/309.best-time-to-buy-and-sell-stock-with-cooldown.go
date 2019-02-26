package main

import "fmt"

/**

题目:
给定一个整数数组，其中第 i 个元素代表了第 i 天的股票价格 。​
设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

示例:
输入: [1,2,3,0,2]
输出: 3
解释: 对应的交易状态为: [买入, 卖出, 冷冻期, 买入, 卖出]

*/

func main() {
	nums := []int{1, 2, 3, 0, 2}
	r := maxProfitWithCooldownOK(nums)
	fmt.Println(r)
}

/**
方法1:
思路同188和123题。
设置三个数组
buy[i]代表截止至i天，买时最大利润
sell[i]代表截止至i天，卖时的最大利润
cool[i]代表截止至i天，冷冻期时候的最大利润

可以得到公式：
buy[i] = max(buy[i-1], cool[i] - price)
sell[i] = max(sell[i-1], buy[i-1] + price)
cool[i] = max(cool[i-1], buy[i-1], sell[i-1])

时间复杂度：O(n)
空间复杂度：O(n)
*/
func maxProfitWithCooldown(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 0 {
		return 0
	}
	buy := make([]int, pricesLen)
	sell := make([]int, pricesLen)
	cool := make([]int, pricesLen)

	minInt := ^int(^uint(0) >> 1)

	for i := 0; i < pricesLen; i++ {
		buy[i] = minInt
		sell[i] = minInt
		cool[i] = minInt
	}

	for i := 0; i < pricesLen; i++ {
		if i == 0 {
			buy[i] = getMaxPrices(buy[i], -prices[i])
			sell[i] = getMaxPrices(sell[i], buy[i]+prices[i])
			cool[i] = getMaxPrices(buy[i], cool[i], sell[i])
		} else {
			buy[i] = getMaxPrices(buy[i-1], cool[i-1]-prices[i])
			sell[i] = getMaxPrices(sell[i-1], buy[i-1]+prices[i])
			cool[i] = getMaxPrices(buy[i-1], cool[i-1], sell[i-1])
		}
	}
	return sell[pricesLen-1]
}

func getMaxPrices(prices ...int) int {
	maxPrice := prices[0]
	for _, v := range prices {
		if v > maxPrice {
			maxPrice = v
		}
	}

	return maxPrice
}

/**
方法2:
可以把问题转化为持有和未持有的问题。
持有的话，肯定是：今天买入，或者前一天未卖出
未持有的话，肯定是：今日卖出，或者前一天卖出今日属于冷冻阶段

时间复杂度：O(n)
空间复杂度：O(1)
*/
func maxProfitWithCooldownOK(prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 {
		return 0
	}

	lastHold := 0
	todayHold := -prices[0]
	lastNoHold := 0
	todayNoHold := 0

	for i := 1; i < pricesLen; i++ {
		lastHold = todayHold
		if i >= 2 {
			todayHold = getMaxV(todayHold, lastNoHold-prices[i])
		} else {
			todayHold = getMaxV(todayHold, -prices[i])
		}

		lastNoHold = todayNoHold
		todayNoHold = getMaxV(todayNoHold, lastHold+prices[i])
	}

	return todayNoHold
}

func getMaxV(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
