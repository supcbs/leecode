package main

import "fmt"

/**
题目：

f(n) = 0 (n=0)
f(n) = 1 (n=1)
f(n) = f(n-1) + f(n-2) (n=2)


*/
func main() {
	f := fiBoNaCciOK(10)
	fmt.Println(f)
}

func fiBoNaCci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fiBoNaCci(n-1) + fiBoNaCci(n-2)
}

func fiBoNaCciOK(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	var dp []int
	dp[0] = 0
	dp[1] = 1

	for i:=2 ; i<=n;i++{
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}