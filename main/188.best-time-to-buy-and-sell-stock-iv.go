package main

import (
	"fmt"
)

/**

题目:
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。
注意: 你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例 1:
输入: [2,4,1], k = 2
输出: 2
解释: 在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。

示例 2:
输入: [3,2,6,5,0,3], k = 2
输出: 7
解释: 在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。
     随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。

*/

func main() {
	nums := []int{3, 2, 6, 5, 0, 3}
	k := 2
	r := maxProfitIV(k, nums)
	fmt.Println(r)
}

/**
方法1:

区分两种情况
1.k>=len/2，那么说明所以涨幅利润都可以吃到，算法同122
2.k<len/2，那么相当于把123的解法给变为n了，还是基于当前买入后最大价和卖出后最大价格来进行处理，最终得到的最大价格就是本题答案


时间复杂度：O(2n)
空间复杂度：O(1)
*/

func maxProfitIV(k int, prices []int) int {
	pricesLen := len(prices)
	if pricesLen <= 1 || k <= 0 {
		return 0
	}

	if k >= pricesLen/2 {
		totalProfit := 0
		for i := 1; i < pricesLen; i++ {
			if prices[i] > prices[i-1] {
				totalProfit += prices[i] - prices[i-1]
			}
		}
		return totalProfit
	} else {
		minInt := ^int(^uint(0) >> 1)
		// 设置最小值
		buyArr := make([]int,k)
		sellArr := make([]int,k)

		for i := 0; i < k; i++ {
			buyArr[i] = minInt
			sellArr[i] = minInt
		}
		for _, price := range prices {
			for i := 0; i < k; i++ {
				if i == 0 {
					buyArr[i] = getMaxIV(buyArr[i], -price)
				} else {
					buyArr[i] = getMaxIV(buyArr[i], sellArr[i-1]-price)
				}
				sellArr[i] = getMaxIV(sellArr[i], buyArr[i]+price)
			}
		}
		return sellArr[k-1]
	}
}


func getMaxIV(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
