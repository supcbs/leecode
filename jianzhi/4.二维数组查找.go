package main

import "fmt"

/**
题目：
一个二维数组，从左到右和从上到下都是递增的。完成一个函数进程数字的查找

分析：
1.可以从右上或者左下进行查找，每次查找会排除掉一行或者一列
（不能是左上或者右下，因为无法排出一列）
*/
func main() {
	arr := [][]int{
		{1,2,8,9},
		{2,4,9,12},
		{4,7,10,13},
		{6,8,11,15},
	}
	n := 7
	r := findInPartialltSortedMatrix(arr, n)
	fmt.Println(r)
}

/**
1.获取左上的坐标点（0，x）
*/
func findInPartialltSortedMatrix(arr [][]int, n int) int {
	if len(arr) <= 0 {
		return -1
	}

	col := 0
	row := len(arr[0])-1
	for col < len(arr) && row > 0 {
		if n == arr[col][row] {
			return 1
		}
		if n > arr[col][row] {
			col ++
		} else {
			row --
		}
	}

	return -1
}