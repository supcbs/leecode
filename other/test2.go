package main

import "fmt"

func main() {
	data := [][]int{
		{300, 500, 560, 400, 160},
		{1000, 100, 200, 340, 690},
		{600, 500, 500, 460, 320},
		{300, 400, 250, 210, 760},
	}

	m := 4
	n := 5

	//fmt.Println(getMaxRed(m, n, 0, 0, data))
	fmt.Println(getMaxRedDP(m, n, data))
}

/**
题目分析：

要想获取到最大的红包路径，一定是前一步较大的金额和当前步相加得到的。
*/

func getMaxRed(m, n, i, j int, data [][]int) int {
	if i >= m || j >= n {
		return 0
	}

	numButton := getMaxRed(m, n, i+1, j, data)
	numRight := getMaxRed(m, n, i, j+1, data)

	if numButton > numRight {
		return data[i][j] + numButton
	}

	return data[i][j] + numRight
}

/**
非递归版的动态规划：
dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + data[i][j]
*/

func getMaxRedDP(m, n int, data [][]int) int {
	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = data[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + data[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + data[i][j]
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j] + data[i][j]
				} else {
					dp[i][j] = dp[i][j-1] + data[i][j]
				}
			}
		}
	}

	fmt.Println(dp)
	return dp[m-1][n-1]
}
