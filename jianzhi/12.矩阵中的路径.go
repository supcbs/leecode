package main

import "fmt"

/**
题目：
给一个布满字母的矩阵，求出是否存在某个单词，只能上下左右的走

[
a,b,t,g,
c,f,c,s,
j,d,e,h
]

寻找bfce

分析：
典型的回溯法进行解题。
1.由于每个字母只能走一遍，所以需要一个数组visited来进行记录走了哪些了。
2.每一步能走的条件是：在范围内 && 值符合当前需要查找的字母 && 没有走过

*/
func main() {
	matrix := [][]string{
		{"a", "b", "t", "g"},
		{"c", "f", "c", "s"},
		{"j", "d", "e", "h"},
	}
	str := "bfce"
	fmt.Println(stringPathInMatrix(matrix, str))
}

func stringPathInMatrix(matrix [][]string, str string) bool {
	if len(matrix) <= 0 || len(str) <= 0 {
		return false
	}

	visited := make([][]bool, 0)
	for i := 0; i < len(matrix); i++ {
		tmp := make([]bool, len(matrix[0]))
		visited = append(visited, tmp)
	}

	fmt.Println(visited)

	pathLength := 0
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if hasPath(matrix, str, pathLength, i, j, len(matrix)-1, len(matrix[0])-1, visited) {
				return true
			}
		}
	}

	return false
}

func hasPath(matrix [][]string, str string, pathLength int, col int, row int, cols int, rows int, visited [][]bool) bool {
	if pathLength >= len(str) {
		return true
	}

	ok := false
	target := fmt.Sprintf("%c", str[pathLength])
	if col >= 0 && row >= 0 && col <= cols && row <= rows && target == matrix[col][row] && !visited[col][row] {
		pathLength++
		visited[col][row] = true
		ok = hasPath(matrix, str, pathLength, col+1, row, cols, rows, visited) ||
			hasPath(matrix, str, pathLength, col-1, row, cols, rows, visited) ||
			hasPath(matrix, str, pathLength, col, row+1, cols, rows, visited) ||
			hasPath(matrix, str, pathLength, col, row-1, cols, rows, visited)

		if !ok {
			pathLength--
			visited[col][row] = false
		}
	}
	return ok
}
