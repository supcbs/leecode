package main

import "fmt"

/**
题目：

本质上也是斐波那契数列

*/
func main() {
	f := fiBoNaCciOK(10)
	fmt.Println(f)
}

func fiBoNaCci(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return fiBoNaCci(n-1) + fiBoNaCci(n-2)
}

func fiBoNaCciOK(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	var dp []int
	dp[1] = 1
	dp[2] = 2

	for i:=3 ; i<=n;i++{
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}