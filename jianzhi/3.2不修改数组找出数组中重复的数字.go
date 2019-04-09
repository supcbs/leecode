package main

import "fmt"

/**
题目：
在不改变数组的情况下有n+1个数组组成的数组，取值范围为1 ～ n，找出重复的数组

分析：
1.可以使用排序+遍历比值的方式，每次比较当前值与前一个是否相等来判断 -- 不行
2.可以采用map的方式，如果重复了就ok了 -- 可以，但是需要额外空间
3.采用规律来进行判断，确认共n个数，那么在1-m中的数字个数如果大于m，那么代表有重复，否则重复的数字在m+1~n之间
*/
func main() {
	arr := []int{2, 3, 5, 4, 1, 2, 6, 7}
	n := 7
	r := duplicationInArrayNotEdit(arr, n)
	fmt.Println(r)
}

/**
1.从0开始遍历，如果当前值等于其下标，则下一个，否则就将当前值放到其下表的位置
2.如果其下标位置的数=当前值，那么结束
*/
func duplicationInArrayNotEdit(arr []int, n int) int {
	if len(arr) <= 0 || n+1 != len(arr) {
		return -1
	}

	start := 0
	end := len(arr) - 1
	for start <= end {
		mid := (start + end) >> 1
		count := calCount(arr, mid)

		if start == end {
			if count > 1 {
				return start
			}
			return -1
		}

		if count > mid {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return -1
}

func calCount(arr []int, num int) int {
	count := 0
	for _, v := range arr {
		if v <= num {
			count++
		}
	}
	return count
}
