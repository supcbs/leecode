package main

import "fmt"

/**
题目：
有m*n的数组，一个机器人从[0,0]开始行走，要求其坐标位数之和不能大于k，求机器人能走的范围格子数。
例如：k=8,[35,37] => 3+5+3+7=18是可以走的

分析：
典型的回溯法进行解题。
1.由于每个字母只能走一遍，所以需要一个数组visited来进行记录走了哪些了。
2.每一步能走的条件是：在范围内 && 值符合当前需要满足小于k && 没有走过
*/
func main() {
	num := robotMove(5, 5, 5)
	fmt.Println(num)
}

func robotMove(k, cols, rows int) int {
	if k <= 0 || rows <= 0 || cols <= 0 {
		return 0
	}

	// 构建已经走过的路
	visited := make([][]bool, 0)
	for i := 0; i < cols; i++ {
		tmp := make([]bool, rows)
		visited = append(visited, tmp)
	}

	fmt.Println(visited)
	count := moveCount(visited, k, cols, rows, 0, 0)

	return count
}

func moveCount(visited [][]bool, k int, cols int, rows int, col int, row int) int {
	count := 0
	if col >= 0 && row >= 0 && col < cols && row < rows && !visited[col][row] && checkNumOK(k, col, row) {
		visited[col][row] = true
		count = 1 + moveCount(visited, k, cols, rows, col+1, row) +
			moveCount(visited, k, cols, rows, col, row+1) +
			moveCount(visited, k, cols, rows, col-1, row) +
			moveCount(visited, k, cols, rows, col, row-1)
	}

	return count
}

func checkNumOK(k, col, row int) bool {
	count := 0
	for col > 0 {
		count += col % 10
		col /= 10
	}

	for row > 0 {
		count += row % 10
		row /= 10
	}

	return count <= k
}
