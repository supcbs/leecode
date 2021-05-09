package main

import "fmt"

/**

题目:
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。

[[1,3,5,7],[10,11,16,20],[23,30,34,60]]
*/

func main() {
	nums := [][]int{
		{1,3,5,7},
		{10,11,16,20},
		{23,30,34,60},
	}
	r := searchMatrix(nums, 3)
	fmt.Println(r)
}

/**
方法1:

时间复杂度：O(n)
空间复杂度：O(1)
*/
func searchMatrix(matrix [][]int, target int) bool {
	// 边界值判断
	if len(matrix) <= 0 {
		return false
	}

	// 初始值设置
	low := 0
	high := len(matrix[0])*len(matrix) - 1
	length := len(matrix[0])

	// 进行二分查找
	for low <= high {
		mid := low + (high - low) / 2
		num := matrix[mid / length][mid % length]
		if num == target {
			return true
		} else if num > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return false
}