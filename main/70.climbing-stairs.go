package main

import "fmt"

/**

题目:
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
注意：给定 n 是一个正整数。

示例 1：
输入： 2
输出： 2
解释： 有两种方法可以爬到楼顶。
1.  1 阶 + 1 阶
2.  2 阶

示例 2：
输入： 3
输出： 3
解释： 有三种方法可以爬到楼顶。
1.  1 阶 + 1 阶 + 1 阶
2.  1 阶 + 2 阶
3.  2 阶 + 1 阶

*/

func main() {
	n := 10
	r := climbStairs(n)
	fmt.Println(r)
}

/**
方法1:

动态规划。
第n = n-1层的方法数目 + n-2层的方法数量

时间复杂度：O(n)
空间复杂度：O(1)
*/
func climbStairs(n int) int {
	if n <= 2{
		return n
	}

	tmp := 1
	res := 2
	for i:=3;i<=n;i++{
		c := res
		res += tmp
		tmp = c
	}

	return res
}

func climbStairs2(n int) int {
	if n <= 2{
		return n
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs3(n int) int {
	if n <= 2{
		return n
	}

	dp := make([]int,n)
	dp[0] = 1
	dp[1] = 2

	for i:=2;i<n;i++{
		dp[i] = dp[i-1]+dp[i-2]
	}

	return dp[n-1]
}


