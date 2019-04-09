package main

import "fmt"

/**

题目:
给定一个三角形，找出自顶向下的最小路径和。每一步只能移动到下一行中相邻的结点上。
例如，给定三角形：

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
自顶向下的最小路径和为 11（即，2 + 3 + 5 + 1 = 11）。

*/

func main() {
	g := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	r := minimumTotalDP(g)
	fmt.Println(r)
}

/**
方法1:

这道题的关键点在于，总尾巴开始进行向上的遍历

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minimumTotal(triangle [][]int) int {
	tLen := len(triangle)
	if tLen <= 0 {
		return 0
	}

	for i := tLen - 1; i >= 0; i-- {
		l := len(triangle[i])
		if l <= 0 {
			continue
		}
		for j := 0; j < l-1; j++ {
			if triangle[i][j] > triangle[i][j+1] {
				triangle[i-1][j] += triangle[i][j+1]
			} else {
				triangle[i-1][j] += triangle[i][j]
			}
		}
	}

	return triangle[0][0]
}

/**
1.最左边的一定只能是当前数+[i-1][0]的数
2.i=j代表最右边，一定是[i-1][j-1]的数
3.中间的数，等于[i-1][j-1]与[i-1][j]直接较小的数
4.遍历最后一排得出最小的数
*/
func minimumTotalDP(triangle [][]int) int {
	tLen := len(triangle)
	if tLen <= 0 {
		return 0
	}

	var dp [][]int
	for i := 0; i < len(triangle); i++ {
		tmp := make([]int, tLen)
		dp = append(dp, tmp)
	}

	for i := 0; i < tLen; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				if i == 0 {
					dp[i][j] = triangle[i][j]
				} else {
					dp[i][j] = triangle[i][j] + dp[i-1][j]
				}
			} else if i == j {
				dp[i][j] = triangle[i][j] + dp[i-1][j-1]
			} else {
				if dp[i-1][j-1] > dp[i-1][j] {
					dp[i][j] = triangle[i][j] + dp[i-1][j]
				} else {
					dp[i][j] = triangle[i][j] + dp[i-1][j-1]
				}
			}
		}
	}

	fmt.Println(dp)
	ret := dp[tLen-1][0]
	for i := 1; i < len(dp[tLen-1]); i++ {
		if ret > dp[tLen-1][i] {
			ret = dp[tLen-1][i]
		}
	}

	return ret
}

func minimumTotalOK(triangle [][]int) int {
	min := 0
	if triangle == nil || len(triangle[0]) == 0{
		return min
	}
	for i:=1; i<len(triangle); i++ {
		length := len(triangle[i])
		triangle[i][0] = triangle[i][0] + triangle[i-1][0]
		triangle[i][length-1] = triangle[i][length-1] + triangle[i-1][length-2]
		for j:=1; j<length-1; j++ {
			if triangle[i-1][j-1] < triangle[i-1][j] {
				triangle[i][j] = triangle[i][j]+triangle[i-1][j-1]
			}else {
				triangle[i][j] = triangle[i][j]+triangle[i-1][j]
			}
		}
	}
	index := len(triangle) - 1
	min = triangle[index][0]

	for i:=1; i<len(triangle[index]); i++ {
		if triangle[index][i] < min {
			min = triangle[index][i]
		}
	}

	return min
}