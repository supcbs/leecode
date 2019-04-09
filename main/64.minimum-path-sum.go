package main

import "fmt"

/**

题目:
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，
使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
示例:
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]

输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。


*/

func main() {
	a := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{2, 4, 1},
	}
	r := minPathSum(a)
	fmt.Println(r)
}

/**
方法1:

dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + dp[i][j]

时间复杂度：O(n)
空间复杂度：O(1)
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		l := len(grid[i])
		for j := 0; j < l; j++ {
			if j == 0 {
				if i == 0 {
					continue
				} else {
					grid[i][j] = grid[i][j] + grid[i-1][j]
				}
			} else if j > 0 {
				if i == 0 {
					grid[i][j] = grid[i][j] + grid[i][j-1]
				} else {
					if 	grid[i-1][j] > grid[i][j-1] {
						grid[i][j] = grid[i][j-1] + grid[i][j]
					} else {
						grid[i][j] = grid[i-1][j] + grid[i][j]
					}
				}
			}
		}
	}

	for _,v:=range grid{
		fmt.Println(v)
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
